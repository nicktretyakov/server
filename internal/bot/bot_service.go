package bot

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"runtime"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"be/internal/datastore/address"
)

type Config struct {
	BotURL                            string
	BookingViewFrontRouteNote string
	RoomViewFrontRouteNote string
}

type ChatService struct {
	cfg   Config
	token string
}

func NewChatService(cfg Config) *ChatService {
	return &ChatService{cfg: cfg}
}

func (s *ChatService) getURL(subPath string) string {
	u, _ := url.Parse(s.cfg.BotURL)
	u.Path = path.Join(u.Path, subPath)

	return u.String()
}

func (s *ChatService) newRequest(method, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(context.Background(), method, s.getURL(path), body)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.token))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getError(body io.Reader) error {
	var pResp ChatAPIError

	if err := json.NewDecoder(body).Decode(&pResp); err != nil {
		return err
	}

	return errors.New(pResp.Message)
}

func (s *ChatService) Login(req *LoginBotRequest) (*LoginBotResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := s.newRequest("POST", "login", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, getError(resp.Body)
	}

	var pResp LoginBotResponse
	if err := json.NewDecoder(resp.Body).Decode(&pResp); err != nil {
		return nil, err
	}

	s.token = pResp.Token

	return &pResp, nil
}

func (s *ChatService) Logout() {
	resp, err := s.newRequest("GET", "logout", nil)
	if err != nil {
		log.Error().Msgf("Logout error: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error().Msgf("Logout error: %v", getError(resp.Body))
	}
}

func (s *ChatService) CreateStream(req *CreateStreamRequest) (*Stream, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := s.newRequest("POST", "stream", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, getError(resp.Body)
	}

	var pResp Stream
	if err := json.NewDecoder(resp.Body).Decode(&pResp); err != nil {
		return nil, err
	}

	return &pResp, nil
}

func (s *ChatService) PublishStream(req *CreateMessageRequest) (*Message, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := s.newRequest("POST", "message", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, getError(resp.Body)
	}

	var pResp Message
	if err := json.NewDecoder(resp.Body).Decode(&pResp); err != nil {
		return nil, err
	}

	return &pResp, nil
}

func (s *ChatService) GetSubscribedStreams() (*ListSubscribesResponse, error) {
	resp, err := s.newRequest("GET", "subscribe", nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, getError(resp.Body)
	}

	var pResp ListSubscribesResponse
	if err = json.NewDecoder(resp.Body).Decode(&pResp); err != nil {
		return nil, err
	}

	return &pResp, nil
}

func (s *ChatService) AddUserToStream(streamID string, userUUIDs []string) error {
	body, err := json.Marshal(userUUIDs)
	if err != nil {
		return err
	}

	resp, err := s.newRequest("POST", fmt.Sprintf("members/%s", streamID), bytes.NewReader(body))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return getError(resp.Body)
	}

	return nil
}

func (s *ChatService) createDynLink(payload string) (string, error) {
	body, err := json.Marshal(struct {
		Payload string
	}{
		Payload: payload,
	})
	if err != nil {
		return "", err
	}

	resp, err := s.newRequest("POST", "/dynlink", bytes.NewReader(body))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", getError(resp.Body)
	}

	var result struct {
		ShortLink string
	}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.ShortLink, nil
}

func (s *ChatService) CreateProposalViewLink(
	id,
	ip,
	secret string,
	addressType address.TypeAddress,
	addressUUID uuid.UUID,
) (string, error) {
	var viewURL string

	if _, err := s.Login(&LoginBotRequest{
		ID: id,
		IP: ip,
		OS: OS{
			Name:    runtime.GOOS,
			Version: runtime.GOARCH,
		},
		Secret: secret,
	}); err != nil {
		return "", err
	}

	defer s.Logout()

	switch addressType {
	case address.BookingAddressType:
		viewURL = s.cfg.BookingViewFrontRouteNote
	case address.RoomAddressType:
		viewURL = s.cfg.RoomViewFrontRouteNote
	case address.UnknownAddressType:
		return "", errors.New("unknown invest object type")
	default:
		return "", errors.New("unknown invest object type")
	}

	viewURL = strings.Replace(viewURL, "[[uuid]]", addressUUID.String(), 1)

	return s.createDynLink(viewURL)
}

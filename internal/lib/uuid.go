package lib

import "github.com/google/uuid"

func UUID() uuid.UUID {
	return uuid.Must(uuid.NewUUID())
}

func ParseUUIDStrings(s []string) ([]uuid.UUID, error) {
	parsedUUIDs := make([]uuid.UUID, 0, len(s))

	for _, _uuid := range s {
		p, err := uuid.Parse(_uuid)
		if err != nil {
			return nil, err
		}

		parsedUUIDs = append(parsedUUIDs, p)
	}

	return parsedUUIDs, nil
}

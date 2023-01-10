package emailnote

//nolint:lll
const HTMLTemplate = `
<html>
<head>
    <title>Технониколь</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body style="padding:0;margin-left:0;margin-top:0;margin-right:0;margin-bottom:0;font-family:Tahoma, Geneva, sans-serif">
<table border="0" align="center" cellpadding="0" cellspacing="0" width="100%"
       style="width:100%;border-spacing:0;padding:0;margin-left:0;margin-top:0;margin-right:0;margin-bottom:0;">
    <tr>
        <td>
            <table border="0" align="center" cellpadding="0" cellspacing="0" width="600"
                   style="width:600px;border-spacing:0;padding:0;margin-left:0;margin-top:0;margin-right:0;margin-bottom:0;">
                <tbody>
                <tr>
                    <td colspan="3" height="11" bgcolor="#ffffff" style="background-color:#ffffff;"></td>
                </tr>
                <tr>
                    <td colspan="1" width="20" bgcolor="#ffffff" style="background-color:#ffffff;"></td>
                    <td colspan="1" width="560" bgcolor="#ffffff" style="background-color:#ffffff;">
                        <table border="0" align="center" cellpadding="0" cellspacing="0" width="100%"
                               style="width:100%;border-spacing:0;padding:0;margin-left:0;margin-top:0;margin-right:0;margin-bottom:0;">
                            <tbody>
                            <tr>
                                <td colspan="1" width="50%" bgcolor="#ffffff" style="background-color:#ffffff;">
                                    <a href="https://www.tn.ru/">
                                        <img alt="Технониколь" width="196" height="35" border="0"
                                             style="width: 196px;height: 35px;"
                                             src="data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjQwIiBoZWlnaHQ9IjQxIiB2aWV3Qm94PSIwIDAgMjQwIDQxIiBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPgo8cGF0aCBkPSJNNDAuNTM5NCAzMi43OTk4TDMyLjQzMjIgNDAuOTk3OUwyLjQ0MTk1IDQxTDAgMzguNTI5NlY4LjE5OTgzTDguMTA4MzcgMEgzOC4wOTc4TDQwLjUzOTQgMi40NjkyMVYzMi43OTk4WiIgZmlsbD0iI0VEMUMyNCIvPgo8cGF0aCBkPSJNMzguMTAwNSAwTDI4Ljk4MjQgOS4yMjE0VjIyLjY4MTJMMjEuNTEzMSAxMi4zMDFINS43NjgzMlYxNS41NjI4SDEwLjQwOTVWMjguMDA1MUwwLjAwMzQxNzk3IDM4LjUyOTZWNDFMMi40NDQ1NSA0MC45OTc5TDEzLjkwNjkgMjkuNDA1N1YxNS41NjI4SDE4LjU0NjVWMjguNjk4NkgyMS45OTk0VjE4LjMxOEwyOS40NjkgMjguNjk4NkgzMi40MzQ4VjEwLjY2ODJMNDAuNTQyNCAyLjQ2OTIxTDQwLjU0MiAwSDM4LjEwMDVaIiBmaWxsPSJ3aGl0ZSIvPgo8cGF0aCBkPSJNNTMuNTEwOCAxNi4yODMySDQ4LjY0NjVWMTIuMzAwNkg2Mi44NzA0VjE2LjI4MzJINTguMDA2NVYyOC42OTgySDUzLjUxMTJWMTYuMjgzMkg1My41MTA4WiIgZmlsbD0iYmxhY2siLz4KPHBhdGggZD0iTTY1LjIzMzQgMTIuMzAxSDc4LjI3NDNWMTYuMTY0Nkg2OS42ODA5VjE4LjY1MDNINzcuNDYzMlYyMi4yMzUzSDY5LjY4MDlWMjQuODM0M0g3OC4zOTA3VjI4LjY5OTFINjUuMjMzNFYxMi4zMDFaIiBmaWxsPSJibGFjayIvPgo8cGF0aCBkPSJNODUuNzg5MiAyMC4zNTk1TDgwLjQ1OTkgMTIuMzAxSDg1LjYyN0w4OC40OTg4IDE2LjkxNTVMOTEuMzk1OSAxMi4zMDFIOTYuNDQ0MUw5MS4xMTgxIDIwLjMxMjhMOTYuNjc3NCAyOC42OTg2SDkxLjUxMDhMODguNDA1MiAyMy43NTY1TDg1LjI3OTIgMjguNjk4Nkg4MC4yMjlMODUuNzg5MiAyMC4zNTk1WiIgZmlsbD0iYmxhY2siLz4KPHBhdGggZD0iTTk4LjgwMjcgMTIuMzAxSDEwMy4yOTdWMTguNDM4N0gxMDkuMDY1VjEyLjMwMUgxMTMuNTU5VjI4LjY5ODZIMTA5LjA2NVYyMi40NjY3SDEwMy4yOTdWMjguNjk4Nkg5OC44MDI3VjEyLjMwMVoiIGZpbGw9ImJsYWNrIi8+CjxwYXRoIGQ9Ik0xMTYuMjU2IDIwLjVDMTE2LjI1NiAxNS43OTAyIDEyMC4wMDkgMTEuOTcyMSAxMjUuMDEzIDExLjk3MjFDMTMwLjAxNiAxMS45NzIxIDEzMy43MjIgMTUuNzQ0MyAxMzMuNzIyIDIwLjVDMTMzLjcyMiAyNS4yMDkxIDEyOS45NjkgMjkuMDI3NiAxMjQuOTY3IDI5LjAyNzZDMTE5Ljk2MSAyOS4wMjggMTE2LjI1NiAyNS4yNTQ2IDExNi4yNTYgMjAuNVpNMTI5LjEzNSAyMC41QzEyOS4xMzUgMTguMTMzMyAxMjcuNDQ0IDE2LjA3MDggMTI0Ljk2NyAxNi4wNzA4QzEyMi41MDkgMTYuMDcwOCAxMjAuODY1IDE4LjA4NTggMTIwLjg2NSAyMC41QzEyMC44NjUgMjIuODY2OCAxMjIuNTU3IDI0LjkyNTYgMTI1LjAxMyAyNC45MjU2QzEyNy40ODkgMjQuOTI2IDEyOS4xMzUgMjIuOTEzIDEyOS4xMzUgMjAuNVoiIGZpbGw9ImJsYWNrIi8+CjxwYXRoIGQ9Ik0xMzYuNDQzIDEyLjMwMUgxNDAuOTRWMTguNDM4N0gxNDYuNzA2VjEyLjMwMUgxNTEuMjAyVjI4LjY5ODZIMTQ2LjcwNlYyMi40NjY3SDE0MC45NFYyOC42OTg2SDEzNi40NDNWMTIuMzAxWiIgZmlsbD0iYmxhY2siLz4KPHBhdGggZD0iTTE1NC42OTggMTIuMzAxSDE1OS4xNDZWMjEuNTI5OUwxNjUuNDI2IDEyLjMwMUgxNjkuNjE3VjI4LjY5ODZIMTY1LjE2OVYxOS40Njk0TDE1OC44OSAyOC42OTg2SDE1NC42OThWMTIuMzAxWiIgZmlsbD0iYmxhY2siLz4KPHBhdGggZD0iTTE3My4xMTQgMTIuMzAxSDE3Ny42MDdWMTguOTU0OUwxODMuMTkgMTIuMzAxSDE4OC41MThMMTgyLjM4IDE5LjM1MTZMMTg4LjcyNiAyOC42OTkxSDE4My4zM0wxNzkuMjUzIDIyLjYwOTNMMTc3LjYwNyAyNC40NTg2VjI4LjY5OTVIMTczLjExNFYxMi4zMDFaIiBmaWxsPSJibGFjayIvPgo8cGF0aCBkPSJNMTg4Ljc2MiAyMC41QzE4OC43NjIgMTUuNzkwMiAxOTIuNTEzIDExLjk3MjEgMTk3LjUxNyAxMS45NzIxQzIwMi41MjEgMTEuOTcyMSAyMDYuMjI4IDE1Ljc0NDMgMjA2LjIyOCAyMC41QzIwNi4yMjggMjUuMjA5MSAyMDIuNDc0IDI5LjAyNzYgMTk3LjQ3MiAyOS4wMjc2QzE5Mi40NjcgMjkuMDI4IDE4OC43NjIgMjUuMjU0NiAxODguNzYyIDIwLjVaTTIwMS42MzkgMjAuNUMyMDEuNjM5IDE4LjEzMzMgMTk5Ljk1MSAxNi4wNzA4IDE5Ny40NzIgMTYuMDcwOEMxOTUuMDE1IDE2LjA3MDggMTkzLjM3MSAxOC4wODU4IDE5My4zNzEgMjAuNUMxOTMuMzcxIDIyLjg2NjggMTk1LjA2MSAyNC45MjU2IDE5Ny41MTcgMjQuOTI1NkMxOTkuOTk2IDI0LjkyNiAyMDEuNjM5IDIyLjkxMyAyMDEuNjM5IDIwLjVaIiBmaWxsPSJibGFjayIvPgo8cGF0aCBkPSJNMjI2LjMwNyAxMi4zMDFIMjMwLjhWMTcuNDUzMUgyMzMuMjU2QzIzNy4xNDggMTcuNDUzMSAyMzkuOTk5IDE5LjMyODkgMjM5Ljk5OSAyMy4wMjc5QzIzOS45OTkgMjYuNTY3OSAyMzcuNDUgMjguNjk5MSAyMzMuNDg5IDI4LjY5OTFIMjI2LjMwN1YxMi4zMDFaTTIzMy4xNDEgMjQuODgxQzIzNC42MjMgMjQuODgxIDIzNS41MDQgMjQuMjAxMiAyMzUuNTA0IDIyLjkxMzFDMjM1LjUwNCAyMS43NjU0IDIzNC42MjMgMjEuMDE1NCAyMzMuMTYzIDIxLjAxNTRIMjMwLjhWMjQuODgxSDIzMy4xNDFaIiBmaWxsPSJibGFjayIvPgo8cGF0aCBkPSJNMjEwLjIwNyAxMi4zMDFWMTkuNDk0NkMyMTAuMjA3IDIxLjE5NDcgMjEwLjAzNyAyMy42NTQgMjA4LjkyNyAyNC41MjM5QzIwOC40MDYgMjQuOTMxNCAyMDcuNTQ3IDI1LjI0MzggMjA3LjAwMiAyNS4yNDM4VjI5LjEwOThIMjA3LjQyOUMyMDguOTgzIDI5LjEwOTggMjEwLjcyNCAyOC40NzggMjExLjM5MyAyOC4wNDg2QzIxNC4xMDMgMjYuMzA4NCAyMTQuNDM4IDIyLjk5ODIgMjE0LjQzOCAxOS4yNTlWMTYuMjgzMkgyMTguMzE3VjI4LjY5ODJIMjIyLjgxMVYxMi4zMDFIMjEwLjIwN1oiIGZpbGw9ImJsYWNrIi8+Cjwvc3ZnPgo=">
                                    </a>
                                </td>
                                <td colspan="1" width="50%" bgcolor="#ffffff" style="background-color:#ffffff;">
                                    <table border="0" align="center" cellpadding="0" cellspacing="0" width="100%"
                                           style="width:100%;border-spacing:0;padding:0;margin-left:0;margin-top:0;margin-right:0;margin-bottom:0;">
                                        <tbody>
                                        <tr>
                                            <td colspan="1" bgcolor="#ffffff" style="background-color:#ffffff;"
                                                align="right">
                                                <p style="margin-top: 0;margin-right: 0;margin-bottom: 0;margin-left: 0;padding-top:0; padding-right: 0;padding-bottom: 0;padding-left: 0;text-decoration:none;text-align: right;font-style: normal;font-weight: normal;font-size: 12px;line-height: 20px;color: #171725;">
                                                    Горячая линия</p>
                                            </td>
                                        </tr>
                                        <tr>
                                            <td colspan="1" bgcolor="#ffffff" style="background-color:#ffffff;"
                                                align="right">
                                                <a href="tel:88002000565"
                                                   style="margin-top: 0;margin-right: 0;margin-bottom: 0;margin-left: 0;padding-top:0; padding-right: 0;padding-bottom: 0;padding-left: 0;text-decoration:none; cursor: pointer;font-style: normal;font-weight: bold;font-size: 17px;line-height: 1.2;color: #ED1C24;">8
                                                    800 200 05 65</a>
                                            </td>
                                        </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                            </tbody>
                        </table>
                    </td>
                    <td colspan="1" width="20" bgcolor="#ffffff" style="background-color:#ffffff;"></td>
                </tr>
                <tr>
                    <td colspan="3" height="12" bgcolor="#ffffff" width="100%" style="background-color:#ffffff;"></td>
                </tr>
                <tr>
                    <td colspan="1" width="20" bgcolor="#fafafa" style="background-color:#fafafa;"></td>
                    <td colspan="1" width="560" bgcolor="#fafafa" style="background-color:#fafafa;">
                        <table border="0" align="center" cellpadding="0" cellspacing="0" width="100%"
                               style="width:100%;border-spacing:0;padding:0;margin-left:0;margin-top:0;margin-right:0;margin-bottom:0;">
                            <tbody>
                            <tr>
                                <td colspan="3" height="31" bgcolor="#fafafa" style="background-color:#fafafa;"></td>
                            </tr>
                            <tr>
                                <td colspan="3" height="40" bgcolor="#fafafa" style="background-color:#fafafa;"
                                    align="left">
                                    <p style="text-align:left;text-decoration:none;margin-top: 0;margin-right: 0;margin-bottom: 0;margin-left: 0;padding-top:0; padding-right: 0;padding-bottom: 0;padding-left: 0;font-style: normal;font-weight: normal;font-size: 20px;line-height: 29px;color: #171725;">
                                        {{range .DataBody}}
                                            {{if .ForReport}}
												<strong> {{.FirstPartBody}} {{.AddressTitle}} (<a href="{{.Link}}" style="text-decoration:none;color:#3b69cd"> {{.Link}} </a>) {{.SecondPartBody}} </strong>
                                            {{else}}
                                                <strong> {{.FirstPartBody}} <a href="{{.Link}}" style="text-decoration:none;color:#3b69cd"> {{.AddressTitle}} </a> {{.SecondPartBody}} </strong>
                                            {{end}}
                                        {{end}}
                                    </p>
                                </td>
                            </tr>
                            <tr>
                                <td colspan="3" height="15" bgcolor="#fafafa" style="background-color:#fafafa;"></td>
                            </tr>
                            </tbody>
                        </table>
                    </td>
                    <td colspan="1" width="20" bgcolor="#fafafa" style="background-color:#fafafa;"></td>
                </tr>
                <tr>
                    <td colspan="3" height="40" bgcolor="#fafafa" style="background-color:#fafafa;"></td>
                </tr>
                </tbody>
            </table>
        </td>
    </tr>
</table>
</body>
</html>
`

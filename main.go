package main

import (
	"send-mail-with-gomail/service"
)

func main() {

	htmlMsg := `<!DOCTYPE html>
<html>
	<head>
		<title> Go Mail </title>
		<meta charset="UTF-8">
		<style>
			table{
				background-color: yellow;
			}
		</style>
	</head>
	<body>
		<img src="https://picsum.photos/id/237/200/300">
		<table>
			<tr> <th> ID </th> <th> Name </th></tr>
			<tr> <td> 1 </td> <td>  Testing </td></tr>
		</table>

	</body>

</html>`

	service.SendEmail([]string{"receiverofthisemail123@gmail.com"}, "sbuject", htmlMsg)

}

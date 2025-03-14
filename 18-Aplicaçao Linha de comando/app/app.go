package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retorna a aplicaçao de linha de comando pronto para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "App linha de comando"
	app.Usage = "Busca IPs e nomes de servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "youtube.com.br",
		}}

	app.Commands = []cli.Command{
		// comandos
		{
			Name:  "ip",
			Usage: "Serve para buscar o IP da internet",
			Flags: flags,

			Action: buscarIPS,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o servidor que o site esta hospedado",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	return app
}

func buscarIPS(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal("erro")
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor)
	}
}

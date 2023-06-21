package freeling

import (
	"github.com/homelight/json"
	"log"
	"strconv"
	"strings"
)

type Client struct {
	sock *socket
}

/*
	Run server using
	> analyze -f es.cfg --flush --output json --server --port 50005
*/

func NewClient() (*Client, error) {
	c, err := newSocket("127.0.0.1", 50005)
	if err != nil {
		return nil, err
	}

	return &Client{sock: c}, nil
}

func (c *Client) Analyze(text string) (*Response, error) {
	log.Println("Analyzing with FreeLing...")
	log.Println(text)

	resp, err := c.sock.Request([]byte(text))
	if err != nil {
		return nil, err
	}

	log.Println("Response:")
	log.Println(string(resp))

	var result Response
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}

	for i, sentence := range result.Sentences {
		result.Sentences[i].ID, _ = strconv.Atoi(sentence.IdStr)
		for j, token := range sentence.Tokens {
			result.Sentences[i].Tokens[j].ID, _ = strconv.Atoi(strings.Split(token.IdStr, ".")[1])
		}
	}

	return &result, err
}

func (c *Client) Close() error {
	return c.sock.Close()
}

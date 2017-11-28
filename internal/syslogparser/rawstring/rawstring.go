package rawstring

import (
	"time"

	"github.com/Cloverhound/go-syslog/internal/syslogparser"
)

type Parser struct {
	buff      []byte
	priority  syslogparser.Priority
	version   int
	content   string
	location  *time.Location
	timestamp time.Time
}

func NewParser(buff []byte) *Parser {
	return &Parser{
		buff:    buff,
		content: string(buff),
	}
}

func (p *Parser) Parse() error {
	p.timestamp = time.Now().Round(time.Second)
	p.version = syslogparser.NO_VERSION

	return nil
}

func (p *Parser) Dump() syslogparser.LogParts {
	return syslogparser.LogParts{
		"timestamp": p.timestamp,
		"content":   p.content,
		"priority":  p.priority.P,
		"facility":  p.priority.F.Value,
		"severity":  p.priority.S.Value,
	}
}

func (p *Parser) Location(location *time.Location) {
	p.location = location
}

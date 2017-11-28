package format

import (
	"bufio"

	"github.com/Cloverhound/go-syslog/internal/syslogparser/rawstring"
)

/* Selecting an 'Automatic' format detects incoming format (i.e. RFC3164 vs RFC5424) and Framing
 * (i.e. RFC6587 s3.4.1 octet counting as described here as RFC6587, and either no framing or
 * RFC6587 s3.4.2 octet stuffing / non-transparent framing, described here as either RFC3164
 * or RFC6587).
 *
 * In essence if you don't know which format to select, or have multiple incoming formats, this
 * is the one to go for. There is a theoretical performance penalty (it has to look at a few bytes
 * at the start of the frame), and a risk that you may parse things you don't want to parse
 * (rogue syslog clients using other formats), so if you can be absolutely sure of your syslog
 * format, it would be best to select it explicitly.
 */

type RawString struct{}

func (f *RawString) GetParser(line []byte) LogParser {
	return &parserWrapper{rawstring.NewParser(line)}
}

func (f *RawString) GetSplitFunc() bufio.SplitFunc {
	return f.rawStringScannerSplit
}

func (f *RawString) rawStringScannerSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return bufio.ScanLines(data, atEOF)
}

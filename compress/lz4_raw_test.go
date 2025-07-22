package compress

import (
	"bytes"
	"testing"

	"github.com/hyperxpizza/parquet-go/parquet"
)

func TestLz4RawCompress(t *testing.T) {
	lz4RawCompressor := compressors[parquet.CompressionCodec_LZ4_RAW]
	input := []byte("Peter Parker")
	compressed := []byte{
		0xc0, 0x50, 0x65, 0x74, 0x65, 0x72, 0x20, 0x50, 0x61, 0x72, 0x6b, 0x65, 0x72,
	}

	// compression
	output := lz4RawCompressor.Compress(input)
	if !bytes.Equal(compressed, output) {
		t.Fatalf("expected output %s but was %s", string(compressed), string(output))
	}

	// uncompression
	output, err := lz4RawCompressor.Uncompress(compressed)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(input, output) {
		t.Fatalf("expected output %s but was %s", string(input), string(output))
	}
}

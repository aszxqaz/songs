package helpers

import (
	"fmt"
	"slices"
	"strings"
)

func ChunkedBySeparator(s string, sep string, chunkSize int) []string {
	splits := strings.Split(s, sep)
	chunksIter := slices.Chunk(splits, chunkSize)
	chunks := []string{}
	for couplet := range chunksIter {
		chunks = append(chunks, strings.Join(couplet, "\n"))
	}
	fmt.Println(chunks)
	return chunks
}

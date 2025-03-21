package controllers

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Variables globales streaming
var (
	segments      []int
	mediaSequence int = 0
	mu            sync.Mutex
)

const (
	segmentDuration     = 10
	totalSegments       = 64
	maxPlaylistSegments = 3
	updateInterval      = segmentDuration * time.Second
)

func init() {
	// Se inicializan los indices
	segments = []int{0, 1, 2}
	go liveStream() // go routine para maneja livestream
}

// Player genera archivo m3u8 con valores indicados
func Player(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Genera el archivo M3U8
	handleM3U8File(w, r, segmentDuration, mediaSequence, segments)
}

// Manejo de actualizacion cada 10 segundos
func liveStream() {
	for {
		fmt.Printf("Segmentos: {%d %d %d} - MediaSequence: %d\n ", segments[0], segments[1], segments[2], mediaSequence) // Log de monitoreo

		time.Sleep(updateInterval)
		mu.Lock()

		mediaSequence++
		nextSegment := (segments[len(segments)-1] + 1) % totalSegments
		segments = append(segments[1:], nextSegment)
		fmt.Println("Actualizacion de livestream!") // Log de monitoreo

		mu.Unlock()
	}
}

func handleM3U8File(w http.ResponseWriter, r *http.Request, targetDuration int, mediaSequence int, segments []int) []byte {

	// Generacion de archivo incremental livestream
	playlistContent := fmt.Sprintf(
		`#EXTM3U
#EXT-X-VERSION:3
#EXT-X-TARGETDURATION:%d
#EXT-X-MEDIA-SEQUENCE:%d
#EXTINF:%d.000000,
segment_%d.ts
#EXTINF:%d.000000,
segment_%d.ts
#EXTINF:%d.000000,
segment_%d.ts`, targetDuration, mediaSequence, targetDuration, segments[0], targetDuration, segments[1], targetDuration, segments[2])

	// Logica retorno archivo incremental
	_, err := w.Write([]byte(playlistContent))
	if err != nil {
		http.Error(w, "Error al generar archivo m3u8", http.StatusInternalServerError)
	}
	return []byte(playlistContent)

}

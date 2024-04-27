package main

import (
	"fmt"
	"sync"
	"time"
)

var max_depth int // Variabel untuk mengatur kedalaman maksimal

func BFS(startUrl, destinationUrl string, single_pathTes bool) Result {
 if startUrl == destinationUrl {
  return Result{
   Paths:        path_found,
   TotalLinks:   0,
   PathLength:   0,
   DurationInMS: 0,
   PathAmount: 0,
  }
 } else {
  fmt.Println("Starting BFS!")
  total_link_visited = 0
  var wait sync.WaitGroup
  var mut sync.Mutex
  var unvisitedQueue []Node // Queue untuk menyimpan link yang belum dikunjungi
  var minPath int // Jalur minimum dari start ke tujuan
  // HAPUS
  single_path = single_pathTes
  // HAPUS
  max_depth = 100

  // start = awal // nanti ini input start
  // destination = akhir // nanti ini input final
  var startNode Node
  // HAPUS
  startNode.link = startUrl // nanti ini input start
  startNode.depth = 0
  destination = destinationUrl // nanti ini input final
  // HAPUS
  unvisitedQueue = append(unvisitedQueue, startNode)

  // Jadikan link pertama sebagai start
  visitedMap := map[string]string{startNode.link: "start"}

  // hitung waktu proses BFS
  begin = time.Now()
  sekarang := begin
  // Loop berhenti jika queue habis atau waktu melebihi 4,5 menit atau kedalaman dengan rute terpendek sudah dicek semua
  for (sekarang.Sub(begin) <= 4*time.Minute+30*time.Second) && len(unvisitedQueue) > 0 && unvisitedQueue[0].depth < max_depth {
   // Cari jumlah proses yang dijalankan (maksimal 200)
   n_child := len(unvisitedQueue)
   if n_child > 350 {
    n_child = 350
   }

   // Lakukan proses BFS
   for i := 0; i < n_child; i++ {
    wait.Add(1)
    go validasiLinkBFS(&unvisitedQueue, visitedMap, &mut, &wait)
   }
   wait.Wait()

   // Bereskan jika hanya mencari 1 path
   if single_path && len(path_found) > 0 {
    path_found = path_found[:1]
    break
   }

   sekarang = time.Now()
  }

  exTime = sekarang.Sub(begin).Milliseconds()
  if len(path_found) > 0 {
   minPath = len(path_found[0])
   for i := 0; i < len(path_found); i++ {
    if len(path_found[i]) < minPath {
     minPath = len(path_found[i])
    }
   }

   var sementara [][]string
   for i := 0; i < len(path_found); i++ {
    if len(path_found[i]) == minPath {
     sementara = append(sementara, path_found[i])
    }
   }
   path_found = sementara

   fmt.Println("List of paths:")
   fmt.Println(path_found)
   fmt.Printf("Found %d path(s), with minimum depth %d\n", len(path_found), minPath-1)
  } else {
   minPath = -1
   fmt.Println("No path found")
  }
  fmt.Println("Exec time:", exTime, "ms")
  fmt.Printf("Link visited: %d\n", total_link_visited)

  pathLength := 0
  if len(path_found) > 0 {
   pathLength = len(path_found[0])
  }

  pathTotal := len(path_found)

  return Result{
   Paths:        path_found,
   TotalLinks:   total_link_visited,
   PathLength:   pathLength,
   DurationInMS: exTime,
   PathAmount:   pathTotal,
  }
 }
}

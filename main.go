package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

const (
    red   = "red"
    green = "green"
    blue  = "blue"
)

var cubesCounts = map[string]int{
    "red":   12,
    "green": 13,
    "blue":  14,
}

type Set struct {
    RedCubes   int
    GreenCubes int
    BlueCubes  int
}

type Game struct {
    GameID        int
    Sets          []Set
    ValidGame     bool
    MaxRedCubes   int
    MaxGreenCubes int
    MaxBlueCubes  int
    SetPower      uint64
}

func main() {
    ptrFile, err := os.Open("example.txt")
    if err != nil {
        log.Fatalf("while trying to open the file %q: %s", ptrFile.Name(), err)
    }
    defer ptrFile.Close()

    sGames := make([]Game, 0)
    scnr := bufio.NewScanner(ptrFile)
    for scnr.Scan() {
        line := scnr.Text()
        gameConf := strings.Split(line, ":")
        game := Game{
            GameID:    -1,
            Sets:      make([]Set, 0),
            ValidGame: true,
        }
        getGameID(gameConf, &game)
        parseGameSets(gameConf, &game)
        sGames = append(sGames, game)
    }

    result := 0
    powerSumResult := uint64(0)
    for _, game := range sGames {
        fmt.Println(game)
        if game.ValidGame {
            result += game.GameID
        }
        powerSumResult += game.SetPower
    }
    fmt.Println(result)
    fmt.Println(powerSumResult)
}

func getGameID(gameConf []string, game *Game) {
    gameInfo := strings.Split(gameConf[0], " ")
    gameID, err := strconv.Atoi(gameInfo[1])
    if err != nil {
        log.Fatalf("while getting the Game ID from the string %q: %s", gameInfo[1], err)
    }
    game.GameID = gameID
}

func parseGameSets(gameConf []string, game *Game) {
    gameSets := strings.Split(gameConf[1], ";")

    for _, gameSet := range gameSets {
        set := Set{}
        cubeCounts := strings.Split(gameSet, ",")
        for _, cubeCount := range cubeCounts {
            cubeCount = strings.TrimSpace(cubeCount)
            switch {
            case strings.Contains(cubeCount, red):
                countInfo := strings.Split(cubeCount, " ")
                count, err := strconv.Atoi(countInfo[0])
                if err != nil {
                    log.Fatalf("while trying to parse cube count from %q: %s", cubeCount, err)
                }
                set.RedCubes = count
                if count > game.MaxRedCubes {
                    game.MaxRedCubes = count
                }
            case strings.Contains(cubeCount, green):
                countInfo := strings.Split(cubeCount, " ")
                count, err := strconv.Atoi(countInfo[0])
                if err != nil {
                    log.Fatalf("while trying to parse cube count from %q: %s", cubeCount, err)
                }
                set.GreenCubes = count
                if count > game.MaxGreenCubes {
                    game.MaxGreenCubes = count
                }
            case strings.Contains(cubeCount, blue):
                countInfo := strings.Split(cubeCount, " ")
                count, err := strconv.Atoi(countInfo[0])
                if err != nil {
                    log.Fatalf("while trying to parse cube count from %q: %s", cubeCount, err)
                }
                set.BlueCubes = count
                if count > game.MaxBlueCubes {
                    game.MaxBlueCubes = count
                }
            }
        }
        game.SetPower = uint64(game.MaxRedCubes) * uint64(game.MaxGreenCubes) * uint64(game.MaxBlueCubes)
        game.Sets = append(game.Sets, set)
        if set.RedCubes > cubesCounts[red] || set.GreenCubes > cubesCounts[green] || set.BlueCubes > cubesCounts[blue] {
            game.ValidGame = false
        }
    }
}

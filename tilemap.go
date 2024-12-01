package main

import (
	"encoding/json"
	"os"
)

func UnmarshalTilemapJSON(data []byte) (TilemapJSON, error) {
	var r TilemapJSON
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TilemapJSON) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TilemapJSON struct {
	Compressionlevel int64                `json:"compressionlevel"`
	Height           int64                `json:"height"`
	Infinite         bool                 `json:"infinite"`
	Layers           []TilemapLayerJSON   `json:"layers"`
	Nextlayerid      int64                `json:"nextlayerid"`
	Nextobjectid     int64                `json:"nextobjectid"`
	Orientation      string               `json:"orientation"`
	Renderorder      string               `json:"renderorder"`
	Tiledversion     string               `json:"tiledversion"`
	Tileheight       int64                `json:"tileheight"`
	Tilesets         []TilemapTilesetJSON `json:"tilesets"`
	Tilewidth        int64                `json:"tilewidth"`
	Type             string               `json:"type"`
	Version          string               `json:"version"`
	Width            int64                `json:"width"`
}

type TilemapLayerJSON struct {
	Data    []int64 `json:"data"`
	Height  int64   `json:"height"`
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	Opacity int64   `json:"opacity"`
	Type    string  `json:"type"`
	Visible bool    `json:"visible"`
	Width   int64   `json:"width"`
	X       int64   `json:"x"`
	Y       int64   `json:"y"`
}

type TilemapTilesetJSON struct {
	Firstgid int64  `json:"firstgid"`
	Source   string `json:"source"`
}

func NewTilemapJSON(filepath string) (*TilemapJSON, error) {
	// Read the file
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into a TilemapJSON struct
	tilemapJSON, err := UnmarshalTilemapJSON(data)
	if err != nil {
		return nil, err
	}

	return &tilemapJSON, nil
}

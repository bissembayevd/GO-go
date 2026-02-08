package storage

import "Goprojects/ASS-1/internal/models"

var Tasks = make(map[int]models.Task)
var NextID = 1

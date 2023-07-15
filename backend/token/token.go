package token

import "time"

const refreshExp = time.Hour * 24 * 90
const refreshNBF = time.Minute * 45
const accessExp = time.Hour

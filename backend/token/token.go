package token

import "time"

const refreshExp = time.Hour * 24 * 90
const refreshNBF = time.Second * 45
const accessExp = time.Second * 60

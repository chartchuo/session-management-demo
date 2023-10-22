package token

import "time"

const refreshExp = time.Hour * 24 * 30
const refreshNBF = time.Second * 45
const accessExp = time.Second * 60

// define now function
var now = time.Now

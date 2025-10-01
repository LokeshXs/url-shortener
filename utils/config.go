package utils

import "time"

const SHORT_CODE_LENGTH = 6;

const DOMAIN = "http://localhost:3000"

const RATE_LIMIT_TIME_WINDOW = time.Minute; // 1 min

const RATE_LIMIT_REQUESTS_WINDOW = 10;
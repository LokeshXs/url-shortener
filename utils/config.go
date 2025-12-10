package utils

import "time"

const SHORT_CODE_LENGTH = 6;



const RATE_LIMIT_TIME_WINDOW = time.Minute; // 1 min

const RATE_LIMIT_REQUESTS_WINDOW = 10;

const SHORTEN_BASE_URL="https://go.urlbit.space"
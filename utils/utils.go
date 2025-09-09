package utils

import "math/rand"


func GenerateShortCode(length int16) string{

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";

	bytesSlice := make([]byte,length);

	for i := range bytesSlice{


		bytesSlice[i] = charset[rand.Intn(len(charset))];
	}

	return string(bytesSlice);
	
}
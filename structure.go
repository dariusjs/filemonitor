
package main

type Config struct {
  Directories []Directory `json:"directories"`
}

type Directory struct {
  Object string `json:"directory"`
}
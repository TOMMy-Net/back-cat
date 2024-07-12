package internal

type SFTP struct {
	Ip       string   `yaml:"sftp-ip" validate:"required"`
	Port     string   `yaml:"sftp-port" validate:"required"`
	SFTPpath string   `yaml:"sftp-path" validate:"required"`
	Settings Settings `yaml:"settings" validate:"required"`
}


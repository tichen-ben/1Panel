package configs

type System struct {
	Port           string `mapstructure:"port"`
	SSL            string `mapstructure:"ssl"`
	DbFile         string `mapstructure:"db_file"`
	DbPath         string `mapstructure:"db_path"`
	LogPath        string `mapstructure:"log_path"`
	DataDir        string `mapstructure:"data_dir"`
	TmpDir         string `mapstructure:"tmp_dir"`
	Cache          string `mapstructure:"cache"`
	Backup         string `mapstructure:"backup"`
	EncryptKey     string `mapstructure:"encrypt_key"`
	BaseDir        string `mapstructure:"base_dir"`
	Mode           string `mapstructure:"mode"`
	RepoUrl        string `mapstructure:"repo_url"`
	Version        string `mapstructure:"version"`
	Username       string `mapstructure:"username"`
	Password       string `mapstructure:"password"`
	Entrance       string `mapstructure:"entrance"`
	IsDemo         bool   `mapstructure:"is_demo"`
	AppRepo        string `mapstructure:"app_repo"`
	ChangeUserInfo bool   `mapstructure:"change_user_info"`
}

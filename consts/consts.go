package consts

const (
	UserHome                    = "/home/blog"                                                                                                                                                                                        // user home directory
	TempDir                     = "/home/blog/tmpdir"                                                                                                                                                                                 // temporary directory
	BlogBackPrefix              = "blog-backup-"                                                                                                                                                                                      // blog back prefix
	StaticPagePackPrefix        = "static-pages-"                                                                                                                                                                                     // static pages pack prefix
	DefaultThemeId              = "caicai_anatole"                                                                                                                                                                                    // default theme name
	BlogVersion                 = "v1.0.0"                                                                                                                                                                                            // version
	FileSeparator               = ""                                                                                                                                                                                                  // path separator
	SuffixFtl                   = ".ftl"                                                                                                                                                                                              // suffix of template file
	MethodKey                   = "method"                                                                                                                                                                                            // custom tag method key
	NeteaseMusicPrefix          = "[music:"                                                                                                                                                                                           // 网易云音乐短代码前缀
	NeteaseMusicIframe          = "<iframe frameborder=\"no\" border=\"0\" marginwidth=\"0\" marginheight=\"0\" width=330 height=86 src=\"//music.163.com/outchain/player?type=2&id=$1&auto=1&height=66\"></iframe>"                  // 网易云音乐 iframe 代码
	NeteaseMusicRegPattern      = "\\[music:(\\d+)\\]"                                                                                                                                                                                // 网易云音乐短代码正则表达式
	BilibiliVideoPrefix         = "[bilibili:"                                                                                                                                                                                        //哔哩哔哩视频短代码前缀
	BilibiliVideoIframe         = "<iframe height=$3 width=$2 src=\"//player.bilibili.com/player.html?aid=$1\" scrolling=\"no\" border=\"0\" frameborder=\"no\" framespacing=\"0\" allowfullscreen=\"true\"> </iframe>"               // 哔哩哔哩视频 iframe 代码
	BilibiliVideoRegPattern     = "\\[bilibili:(\\d+)\\,(\\d+)\\,(\\d+)\\]"                                                                                                                                                           // 哔哩哔哩视频正则表达式
	YoutubeVideoPrefix          = "[youtube:"                                                                                                                                                                                         // YouTube 视频短代码前缀
	YouTubeVideoIframe          = "<iframe width=$2 height=$3 src=\"https://www.youtube.com/embed/$1\" frameborder=\"0\" allow=\"accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture\" allowfullscreen></iframe>" // YouTube 视频 iframe 代码
	YouTubeRegPattern           = "\\[youtube:(\\w+)\\,(\\d+)\\,(\\d+)\\]"                                                                                                                                                            //YouTube 视频正则表达式
	BlogAdminReleasesLatest     = "https://api.github.com/repos/halo-dev/halo-admin/releases/latest"                                                                                                                                  // github api url for blog release
	BlogAdminVersionRegex       = "blog-admin-\\d+\\.\\d+(\\.\\d+)?(-\\S*)?\\.zip"                                                                                                                                                    // blog admin version regex
	BlogAdminRelativePath       = "templates/admin/"
	BlogAdminRelativeBackupPath = "templates/admin-backup/"
	ApiAccessKeyHeaderName      = "API-Authorization"   // content token header name
	AdminTokenHeaderName        = "ADMIN-Authorization" // admin token header name
	AdminTokenQueryName         = "admin_token"         // admin token param name
	TempToken                   = "temp_token"          // temporary token
	ApiAccessKeyQueryName       = "api_access_key"      // content api token param name
	OneTimeTokenQueryName       = "ott"
	OneTimeTokenHeaderName      = "ott"
	UserSessionKey              = "user_session"
)

const (
	IsInstalled = false          // is blog installed
	Theme       = DefaultThemeId // current activied theme
	Birthday    = ""             // blog birthday
	DevMode     = false          // developer mode
)

// email
const (
	EmailHost     = ""     // email sender host
	EmailProtocol = "smtp" // email sender protocol
	EmailSSLPort  = "465"  // SSL port
	EmailUsername = ""     // email sender username
	EmailPassword = ""     // email sender password
	EmailFromName = ""     // email sender name
	EmailEnabled  = "fase" // is enabled email sender
)

// post
const (
	PostStatusPublished = iota // published
	PostStatusDraft            // draft
	PostStatusRecycle          // recycle
	PostStatusIntimate         // intimate
)

// comment
const (
	CommentStatusPublished = iota // puslished
	CommentStatusAuditing         // auditing
	CommentStatusRecycle          //recycle
)

// journal
const (
	JournalStatusPublished = iota // puslished
	JournalStatusAuditing         // auditing
	JournalStatusRecycle          //recycle
)

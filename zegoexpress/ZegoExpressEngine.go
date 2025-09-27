package zegoexpress

type ZegoRoomMode int

const (
	ZegoRoomModeSingleRoom ZegoRoomMode = iota
	ZegoRoomModeMultiRoom
)

type ZegoScenario int

const (
  ZegoScenarioDefault ZegoScenario = 3 + iota
  ZegoScenarioStandardVideoCall
  ZegoScenarioHighQualityVideoCall
  ZegoScenarioStandardChatroom
  ZegoScenarioHighQualityChatroom
  ZegoScenarioBroadcast
  ZegoScenarioKaraoke
  ZegoScenarioStandardVoiceCall
)

type ZegoRoomState int

const (
	ZegoRoomStateDisconnected ZegoRoomState = iota
	ZegoRoomStateConnecting
	ZegoRoomStateConnected
)

type ZegoRoomStateChangedReason int

const (
	ZegoRoomStateChangedReasonLogining ZegoRoomStateChangedReason = iota
	ZegoRoomStateChangedReasonLogined
	ZegoRoomStateChangedReasonLoginFailed
	ZegoRoomStateChangedReasonReconnecting
	ZegoRoomStateChangedReasonReconnected
	ZegoRoomStateChangedReasonReconnectFailed
	ZegoRoomStateChangedReasonKickOut
	ZegoRoomStateChangedReasonLogout
	ZegoRoomStateChangedReasonLogoutFailed
)

type ZegoRoomStreamListType int

const (
	ZegoRoomStreamListTypePlay ZegoRoomStreamListType = iota
	ZegoRoomStreamListTypeAll
)

type ZegoPublishChannel int

const (
	ZegoPublishChannelMain ZegoPublishChannel = iota
	ZegoPublishChannelAux
	ZegoPublishChannelThird
	ZegoPublishChannelFourth
)

type ZegoPublisherState int

const (
	ZegoPublisherStateNoPublish ZegoPublisherState = iota
	ZegoPublisherStatePublishRequesting
	ZegoPublisherStatePublishing
)

type ZegoStreamCensorshipMode int

const (
	ZegoStreamCensorshipModeNone ZegoStreamCensorshipMode = iota
	ZegoStreamCensorshipModeAudio
	ZegoStreamCensorshipModeVideo
	ZegoStreamCensorshipModeAudioAndVideo
)

type ZegoCapabilityNegotiationType int

const (
	ZegoCapabilityNegotiationTypeNone ZegoCapabilityNegotiationType = iota
	ZegoCapabilityNegotiationTypeAll
	ZegoCapabilityNegotiationTypePublisher
)

type ZegoUpdateType int

const (
	ZegoUpdateTypeAdd ZegoUpdateType = iota
	ZegoUpdateTypeDelete
)

type ZegoStreamQualityLevel int

const (
	ZegoStreamQualityLevelExcellent ZegoStreamQualityLevel = iota
	ZegoStreamQualityLevelGood
	ZegoStreamQualityLevelMedium
	ZegoStreamQualityLevelBad
	ZegoStreamQualityLevelDie
	ZegoStreamQualityLevelUnknown
)

type ZegoVideoCodecID int

const (
	ZegoVideoCodecIDDefault ZegoVideoCodecID = iota
	ZegoVideoCodecIDSVC
	ZegoVideoCodecIDVP8
	ZegoVideoCodecIDH265
	ZegoVideoCodecIDH264DualStream
	ZegoVideoCodecIDUnknown ZegoVideoCodecID = 100
)

type ZegoAudioChannel int

const (
	ZegoAudioChannelUnknown ZegoAudioChannel = iota
	ZegoAudioChannelMono
	ZegoAudioChannelStereo
)

type ZegoAudioCodecID int

const (
	ZegoAudioCodecIDDefault ZegoAudioCodecID = iota
	ZegoAudioCodecIDNormal
	ZegoAudioCodecIDNormal2
	ZegoAudioCodecIDNormal3
	ZegoAudioCodecIDLow
	ZegoAudioCodecIDLow2
	ZegoAudioCodecIDLow3
)

type ZegoAudioSourceType int

const (
	ZegoAudioSourceTypeDefault ZegoAudioSourceType = iota
	ZegoAudioSourceTypeCustom
	ZegoAudioSourceTypeMediaPlayer
	ZegoAudioSourceTypeNone
	ZegoAudioSourceTypeMicrophone
	ZegoAudioSourceTypeMainPublishChannel
)

type ZegoAudioSampleRate int

const (
	ZegoAudioSampleRateUnknown ZegoAudioSampleRate = 0
	ZegoAudioSampleRate8K ZegoAudioSampleRate = 8000
	ZegoAudioSampleRate16K ZegoAudioSampleRate = 16000
	ZegoAudioSampleRate22K ZegoAudioSampleRate = 22050
	ZegoAudioSampleRate24K ZegoAudioSampleRate = 24000
	ZegoAudioSampleRate32K ZegoAudioSampleRate = 32000
	ZegoAudioSampleRate44K ZegoAudioSampleRate = 44000
	ZegoAudioSampleRate48K ZegoAudioSampleRate = 48000
)

type ZegoStreamEvent int

const (
	ZegoStreamEventPublishStart        ZegoStreamEvent = 100 + iota // 100
	ZegoStreamEventPublishSuccess                                   // 101
	ZegoStreamEventPublishFail                                      // 102
	ZegoStreamEventRetryPublishStart                                // 103
	ZegoStreamEventRetryPublishSuccess                              // 104
	ZegoStreamEventRetryPublishFail                                 // 105
	ZegoStreamEventPublishEnd                                       // 106

	ZegoStreamEventPlayStart        ZegoStreamEvent = 200 + iota // 200
	ZegoStreamEventPlaySuccess                                   // 201
	ZegoStreamEventPlayFail                                      // 202
	ZegoStreamEventRetryPlayStart                                // 203
	ZegoStreamEventRetryPlaySuccess                              // 204
	ZegoStreamEventRetryPlayFail                                 // 205
	ZegoStreamEventPlayEnd                                       // 206
)

type ZegoPlayerState int

const (
	ZegoPlayerStateNoPlay ZegoPlayerState = iota
	ZegoPlayerStatePlayRequesting
	ZegoPlayerStatePlaying
)

type ZegoEngineProfile struct {
  AppID uint32
  AppSign string
  Scenario ZegoScenario
}

type ZegoEngineConfig struct {
	LogConfig *ZegoLogConfig
	AdvancedConfig map[string]string
}

type ZegoLogConfig struct {
	LogPath string
	LogSize uint64
	LogCount uint32
}

func NewZegoLogConfig() *ZegoLogConfig {
	return &ZegoLogConfig{
		LogSize: 5 * 1024 * 1024,
		LogCount: 3,
	}
}

type ZegoUser struct {
	UserID   string
	UserName string
}

type ZegoRoomConfig struct {
	MaxMemberCount uint32
	IsUserStatusNotify bool
	Token string
	CapabilityNegotiationTypes uint32
	RoomType uint32
}

type ZegoStream struct {
	User      ZegoUser
	StreamID  string
	ExtraInfo string
}

type ZegoPublisherConfig struct {
	RoomID string
	ForceSynchronousNetworkTime int
	StreamCensorshipMode ZegoStreamCensorshipMode
	StreamCensorFlag int
	CodecNegotiationType ZegoCapabilityNegotiationType
	StreamTitle string
}

type ZegoAudioConfig struct {
	Bitrate int
	Channel ZegoAudioChannel
	CodecID ZegoAudioCodecID
}

type ZegoCustomAudioConfig struct {
	SourceType ZegoAudioSourceType
}

type ZegoAudioFrameParam struct {
	SampleRate ZegoAudioSampleRate
	Channel ZegoAudioChannel
}

type ZegoPublishStreamQuality struct {
	VideoCaptureFPS  float64
	VideoEncodeFPS   float64
	VideoSendFPS     float64
	VideoKBPS        float64
	AudioCaptureFPS  float64
	AudioSendFPS     float64
	AudioKBPS        float64
	Rtt              int
	PacketLostRate   float64
	Level            ZegoStreamQualityLevel
	IsHardwareEncode bool
	VideoCodecID     ZegoVideoCodecID
	TotalSendBytes   float64
	AudioSendBytes   float64
	VideoSendBytes   float64
}

type ZegoPlayStreamQuality struct {
	VideoRecvFPS              float64
	VideoDejitterFPS          float64
	VideoDecodeFPS            float64
	VideoRenderFPS            float64
	VideoKBPS                 float64
	VideoBreakRate            float64
	AudioRecvFPS              float64
	AudioDejitterFPS          float64
	AudioDecodeFPS            float64
	AudioRenderFPS            float64
	AudioKBPS                 float64
	AudioBreakRate            float64
	Mos                       float64
	Rtt                       int
	PacketLostRate            float64
	PeerToPeerDelay           int
	PeerToPeerPacketLostRate  float64
	Level                     ZegoStreamQualityLevel
	Delay                     int
	AvTimestampDiff           int
	IsHardwareDecode          bool
	VideoCodecID              ZegoVideoCodecID
	TotalRecvBytes            float64
	AudioRecvBytes            float64
	VideoRecvBytes            float64
	AudioCumulativeBreakCount uint32
	AudioCumulativeBreakTime  uint32
	AudioCumulativeBreakRate  float64
	AudioCumulativeDecodeTime uint32
	VideoCumulativeBreakCount uint32
	VideoCumulativeBreakTime  uint32
	VideoCumulativeBreakRate  float64
	VideoCumulativeDecodeTime uint32
	MuteVideo                 int
	MuteAudio                 int
}

type ZegoMediaSideInfo struct {
  StreamID string
  SEIData []uint8
  TimestampNs int64
  ModuleType int
}

type ZegoRoomStreamList struct {
	PublishStreamList []ZegoStream
	PlayStreamList []ZegoStream
}

type IZegoEventHandler interface {
	OnDebugError(errorCode int, funcName string, info string)

	OnRoomStateUpdate(roomID string, state ZegoRoomState, errorCode int, extendedData string)
	OnRoomStreamUpdate(roomID string, updateType ZegoUpdateType, streamList []ZegoStream, extendedData string)
	OnRoomStateChanged(roomID string, reason ZegoRoomStateChangedReason, errorCode int, extendedData string)
	OnRoomTokenWillExpire(roomID string, remainTimeInSecond int)

	OnPublisherStateUpdate(streamID string, state ZegoPublisherState, errorCode int, extendedData string)
	OnPublisherQualityUpdate(streamID string, quality ZegoPublishStreamQuality)
	OnPublisherStreamEvent(eventID ZegoStreamEvent, streamID string, extendedData string)

	OnPlayerStateUpdate(streamID string, state ZegoPlayerState, errorCode int, extendedData string)
	OnPlayerQualityUpdate(streamID string, quality ZegoPlayStreamQuality)
  OnPlayerRecvMediaSideInfo(info ZegoMediaSideInfo)
  OnPlayerStreamEvent(eventID ZegoStreamEvent, streamID string, extendedData string)
}

type IZegoAudioDataHandler interface {
	OnPlayerAudioData(data []uint8, param ZegoAudioFrameParam, streamID string)
}

type IZegoApiCalledEventHandler interface {
	OnApiCalledResult(errorCode int, funcName string, info string)
}

type ZegoRoomLoginCallback func(errorCode int, extendedData string)
type ZegoRoomLogoutCallback func(errorCode int, extendedData string)
type ZegoIMSendBroadcastMessageCallback func(errorCode int, messageID uint64)

type IZegoExpressEngine interface {
	EnableDebugAssistant(enable bool)
	LoginRoom(roomID string, user ZegoUser, config *ZegoRoomConfig, callback ZegoRoomLoginCallback)
	LogoutRoom(roomID string, callback ZegoRoomLogoutCallback)
	RenewToken(roomID string, token string)
	SendBroadcastMessage(roomID string, message string, callback ZegoIMSendBroadcastMessageCallback)
	GetRoomStreamList(roomID string, streamListType ZegoRoomStreamListType) ZegoRoomStreamList

	StartPublishingStream(streamID string, config ZegoPublisherConfig, channel ZegoPublishChannel)
	StopPublishingStream(channel ZegoPublishChannel)
	SetAudioConfig(config ZegoAudioConfig, channel ZegoPublishChannel)
	EnableAEC(enable bool)
	EnableCustomAudioIO(enable bool, config *ZegoCustomAudioConfig, channel ZegoPublishChannel)
	SendSEI(data []uint8, channel ZegoPublishChannel)

	SetAudioDataHandler(handler IZegoAudioDataHandler)
	StartAudioDataObserver(observerBitMask uint32, param ZegoAudioFrameParam)
	StopAudioDataObserver()

	StartPlayingStream(streamID string)
	StopPlayingStream(streamID string)

	SendCustomAudioCapturePCMData(data []uint8, param ZegoAudioFrameParam, channel ZegoPublishChannel)
	FetchCustomAudioRenderPCMData(data []uint8, param ZegoAudioFrameParam)
}

func CreateEngine(profile ZegoEngineProfile, eventHandler *IZegoEventHandler) *IZegoExpressEngine {
	return createEngine(profile, eventHandler)
}

type ZegoDestroyCompletionCallback func()
func DestroyEngine(engine *IZegoExpressEngine, callback ZegoDestroyCompletionCallback) {
	destroyEngine(engine, callback)
}

func GetEngine() *IZegoExpressEngine {
	return getEngine()
}

func SetEngineConfig(config ZegoEngineConfig) {
	setEngineConfig(config)
}

func SetLogConfig(config ZegoLogConfig) {
	setLogConfig(config)
}

func SetRoomMode(mode ZegoRoomMode) {
	setRoomMode(mode)
}

func GetVersion() string {
	return getVersion()
}

func SetApiCalledCallback(callback *IZegoApiCalledEventHandler) {
	setApiCalledCallback(callback)
}
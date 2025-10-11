package zegoexpress

// Room mode.
type ZegoRoomMode int

const (
	// Single room mode.
	ZegoRoomModeSingleRoom ZegoRoomMode = iota

	// Multiple room mode.
	ZegoRoomModeMultiRoom
)

// Room scenario.
type ZegoScenario int

const (
	// Available since: 3.0.0. Description: The default (generic) scenario. If none of the following scenarios conform to your actual application scenario, this default scenario can be used.
	ZegoScenarioDefault ZegoScenario = 3 + iota

	// Available since: 3.0.0. Description: Standard video call scenario, it is suitable for one-to-one video call scenarios.
	ZegoScenarioStandardVideoCall

	// Available since: 3.0.0. Description: High quality video call scenario, it is similar to the standard video call scenario, but this scenario uses a higher video frame rate, bit rate, and resolution (540p) by default, which is suitable for video call scenario with high image quality requirements.
	ZegoScenarioHighQualityVideoCall

	// Available since: 3.0.0. Description: Standard chatroom scenario, suitable for multi-person pure voice calls (low data usage). Note: On the ExpressVideo SDK, the camera is not enabled by default in this scenario.
	ZegoScenarioStandardChatroom

	// Available since: 3.0.0. Description: High quality chatroom scenario, it is similar to the standard chatroom scenario, but this scenario uses a higher audio bit rate than the standard chatroom scenario by default. It is suitable for multi-person pure voice call scenarios with high requirements on sound quality. Note: On the ExpressVideo SDK, the camera is not enabled by default in this scenario.
	ZegoScenarioHighQualityChatroom

	// Available since: 3.0.0. Description: Live broadcast scenario, it is suitable for one-to-many live broadcast scenarios such as shows, games, e-commerce, and large educational classes. The audio and video quality, fluency, and compatibility have been optimized. Note: Even in live broadcast scenarios, the SDK has no business "roles" (such as anchors and viewers), and all users in the room can publish and play streams.
	ZegoScenarioBroadcast

	// Available since: 3.0.0. Description: Karaoke (KTV) scenario, it is suitable for real-time chorus and online karaoke scenarios, and has optimized delay, sound quality, ear return, echo cancellation, etc., and also ensures accurate alignment and ultra-low delay when multiple people chorus.
	ZegoScenarioKaraoke

	// Available since: 3.3.0. Description: Standard voice call scenario, it is suitable for one-to-one video or voice call scenarios. Note: On the ExpressVideo SDK, the camera is not enabled by default in this scenario.
	ZegoScenarioStandardVoiceCall
)

// Room state.
type ZegoRoomState int

const (
	// Unconnected state, enter this state before logging in and after exiting the room. If there is a steady state abnormality in the process of logging in to the room, such as AppID or Token are incorrect, or if the same user name is logged in elsewhere and the local end is KickOut, it will enter this state.
	ZegoRoomStateDisconnected ZegoRoomState = iota

	// The state that the connection is being requested. It will enter this state after successful execution login room function. The display of the UI is usually performed using this state. If the connection is interrupted due to poor network quality, the SDK will perform an internal retry and will return to the requesting connection status.
	ZegoRoomStateConnecting

	// The status that is successfully connected. Entering this status indicates that the login to the room has been successful. The user can receive the callback notification of the user and the stream information in the room.
	ZegoRoomStateConnected
)

// Room state change reason.
type ZegoRoomStateChangedReason int

const (
	// Logging in to the room. When calling [loginRoom] to log in to the room or [switchRoom] to switch to the target room, it will enter this state, indicating that it is requesting to connect to the server. The application interface is usually displayed through this state.
	ZegoRoomStateChangedReasonLogining ZegoRoomStateChangedReason = iota

	// Log in to the room successfully. When the room is successfully logged in or switched, it will enter this state, indicating that the login to the room has been successful, and users can normally receive callback notifications of other users in the room and all stream information additions and deletions.
	ZegoRoomStateChangedReasonLogined

	// Failed to log in to the room. When the login or switch room fails, it will enter this state, indicating that the login or switch room has failed, for example, AppID or Token is incorrect, etc.
	ZegoRoomStateChangedReasonLoginFailed

	// The room connection is temporarily interrupted. If the interruption occurs due to poor network quality, the SDK will retry internally.
	ZegoRoomStateChangedReasonReconnecting

	// The room is successfully reconnected. If there is an interruption due to poor network quality, the SDK will retry internally, and enter this state after successful reconnection.
	ZegoRoomStateChangedReasonReconnected

	// The room fails to reconnect. If there is an interruption due to poor network quality, the SDK will retry internally, and enter this state after the reconnection fails.
	ZegoRoomStateChangedReasonReconnectFailed

	// Kicked out of the room by the server. For example, if you log in to the room with the same user name in other places, and the local end is kicked out of the room, it will enter this state.
	ZegoRoomStateChangedReasonKickOut

	// Logout of the room is successful. It is in this state by default before logging into the room. When calling [logoutRoom] to log out of the room successfully or [switchRoom] to log out of the current room successfully, it will enter this state.
	ZegoRoomStateChangedReasonLogout

	// Failed to log out of the room. Enter this state when calling [logoutRoom] fails to log out of the room or [switchRoom] fails to log out of the current room internally.
	ZegoRoomStateChangedReasonLogoutFailed
)

// Get room stream list type.
type ZegoRoomStreamListType int

const (
	// List of all online streams in the current room, excluding your own streams
	ZegoRoomStreamListTypePlay ZegoRoomStreamListType = iota

	// List of all online streams in the current room, including your own streams
	ZegoRoomStreamListTypeAll
)

// Publish channel.
type ZegoPublishChannel int

const (
	// The main (default/first) publish channel.
	ZegoPublishChannelMain ZegoPublishChannel = iota

	// The auxiliary (second) publish channel
	ZegoPublishChannelAux

	// The third publish channel
	ZegoPublishChannelThird

	// The fourth publish channel
	ZegoPublishChannelFourth
)

// Publish stream status.
type ZegoPublisherState int

const (
	// The state is not published, and it is in this state before publishing the stream. If a steady-state exception occurs in the publish process, such as AppID or Token are incorrect, or if other users are already publishing the stream, there will be a failure and enter this state.
	ZegoPublisherStateNoPublish ZegoPublisherState = iota

	// The state that it is requesting to publish the stream after the [startPublishingStream] function is successfully called. The UI is usually displayed through this state. If the connection is interrupted due to poor network quality, the SDK will perform an internal retry and will return to the requesting state.
	ZegoPublisherStatePublishRequesting

	// The state that the stream is being published, entering the state indicates that the stream has been successfully published, and the user can communicate normally.
	ZegoPublisherStatePublishing
)

// Publish CensorshipMode.
type ZegoStreamCensorshipMode int

const (
	// no censorship.
	ZegoStreamCensorshipModeNone ZegoStreamCensorshipMode = iota

	// only censorship stream audio.
	ZegoStreamCensorshipModeAudio

	// only censorship stream video.
	ZegoStreamCensorshipModeVideo

	// censorship stream audio and video.
	ZegoStreamCensorshipModeAudioAndVideo
)

// Type of capability negotiation for publish stream references.
type ZegoCapabilityNegotiationType int

const (
	// no reference to the outcome of the capability negotiation.
	ZegoCapabilityNegotiationTypeNone ZegoCapabilityNegotiationType = iota

	// refer to the outcome of the capability negotiation of all user in the room.
	ZegoCapabilityNegotiationTypeAll

	// refer to the outcome of the capability negotiation of publisher in the room.
	ZegoCapabilityNegotiationTypePublisher
)

// Update type.
type ZegoUpdateType int

const (
	// Add
	ZegoUpdateTypeAdd ZegoUpdateType = iota

	// Delete
	ZegoUpdateTypeDelete
)

// Stream quality level.
type ZegoStreamQualityLevel int

const (
	// Excellent
	ZegoStreamQualityLevelExcellent ZegoStreamQualityLevel = iota

	// Good
	ZegoStreamQualityLevelGood

	// Normal
	ZegoStreamQualityLevelMedium

	// Bad
	ZegoStreamQualityLevelBad

	// Failed
	ZegoStreamQualityLevelDie

	// Unknown
	ZegoStreamQualityLevelUnknown
)

// Video codec ID.
type ZegoVideoCodecID int

const (
	// Default (H.264)
	ZegoVideoCodecIDDefault ZegoVideoCodecID = iota

	// Scalable Video Coding (H.264 SVC)
	ZegoVideoCodecIDSVC

	// VP8
	ZegoVideoCodecIDVP8

	// H.265
	ZegoVideoCodecIDH265

	// Dualstream Scalable Video Coding
	ZegoVideoCodecIDH264DualStream

	// Unknown Video Coding
	ZegoVideoCodecIDUnknown ZegoVideoCodecID = 100
)

// Audio channel type.
type ZegoAudioChannel int

const (
	// Unknown
	ZegoAudioChannelUnknown ZegoAudioChannel = iota

	// Mono
	ZegoAudioChannelMono

	// Stereo
	ZegoAudioChannelStereo
)

// Audio codec ID.
type ZegoAudioCodecID int

const (
	// Default, determined by the [scenario] when calling [createEngine].
	ZegoAudioCodecIDDefault ZegoAudioCodecID = iota

	// Can be used for RTC and CDN streaming; bitrate range from 10kbps to 128kbps; supports stereo; latency is around 500ms. Server cloud transcoding is required when communicating with the Web SDK, and it is not required when relaying to CDN.
	ZegoAudioCodecIDNormal

	// Can be used for RTC and CDN streaming; good compatibility; bitrate range from 16kbps to 192kbps; supports stereo; latency is around 350ms; the sound quality is worse than [Normal] in the same (low) bitrate. Server cloud transcoding is required when communicating with the Web SDK, and it is not required when relaying to CDN.
	ZegoAudioCodecIDNormal2

	// Not recommended; if you need to use it, please contact ZEGO technical support. Can only be used for RTC streaming.
	ZegoAudioCodecIDNormal3

	// Not recommended; if you need to use it, please contact ZEGO technical support. Can only be used for RTC streaming.
	ZegoAudioCodecIDLow

	// Not recommended; if you need to use it, please contact ZEGO technical support. Can only be used for RTC streaming; maximum bitrate is 16kbps.
	ZegoAudioCodecIDLow2

	// Can only be used for RTC streaming; bitrate range from 6kbps to 192kbps; supports stereo; latency is around 200ms; Under the same bitrate (low bitrate), the sound quality is significantly better than [Normal] and [Normal2]; low CPU overhead. Server cloud transcoding is not required when communicating with the Web SDK, and it is required when relaying to CDN.
	ZegoAudioCodecIDLow3
)

// Audio capture source type.
type ZegoAudioSourceType int

const (
	// Default audio capture source (the main channel uses custom audio capture by default; the aux channel uses the same sound as main channel by default).
	ZegoAudioSourceTypeDefault ZegoAudioSourceType = iota

	// Use custom audio capture, refer to [enableCustomAudioIO] or [setAudioSource].
	ZegoAudioSourceTypeCustom

	// Use media player as audio source, only support aux channel.
	ZegoAudioSourceTypeMediaPlayer

	// No audio source. This audio source type can only be used in [setAudioSource] interface, has no effect when used in [enableCustomAudioIO] interface.
	ZegoAudioSourceTypeNone

	// Using microphone as audio source. This audio source type can only be used in [setAudioSource] interface, has no effect when used in [enableCustomAudioIO] interface.
	ZegoAudioSourceTypeMicrophone

	// Using main channel as audio source. Ineffective when used in main channel. This audio source type can only be used in [setAudioSource] interface, has no effect when used in [enableCustomAudioIO] interface.
	ZegoAudioSourceTypeMainPublishChannel
)

// audio sample rate.
type ZegoAudioSampleRate int

const (
	// Unknown
	ZegoAudioSampleRateUnknown ZegoAudioSampleRate = 0

	// 8K
	ZegoAudioSampleRate8K ZegoAudioSampleRate = 8000

	// 16K
	ZegoAudioSampleRate16K ZegoAudioSampleRate = 16000

	// 22.05K
	ZegoAudioSampleRate22K ZegoAudioSampleRate = 22050

	// 24K
	ZegoAudioSampleRate24K ZegoAudioSampleRate = 24000

	// 32K
	ZegoAudioSampleRate32K ZegoAudioSampleRate = 32000

	// 44.1K
	ZegoAudioSampleRate44K ZegoAudioSampleRate = 44100

	// 48K
	ZegoAudioSampleRate48K ZegoAudioSampleRate = 48000
)

// Publish or play stream event
type ZegoStreamEvent int

const (
	// Start publishing stream
	ZegoStreamEventPublishStart ZegoStreamEvent = 100 + iota

	// The first publish stream was successful
	ZegoStreamEventPublishSuccess

	// Failed to publish stream for the first time
	ZegoStreamEventPublishFail

	// Start retrying publishing stream
	ZegoStreamEventRetryPublishStart

	// Retry publishing stream successfully
	ZegoStreamEventRetryPublishSuccess

	// Failed to retry publishing stream
	ZegoStreamEventRetryPublishFail

	// End of publishing stream
	ZegoStreamEventPublishEnd
)

const (
	// Start playing stream
	ZegoStreamEventPlayStart ZegoStreamEvent = 200 + iota

	// The first play stream was successful
	ZegoStreamEventPlaySuccess

	// Failed to play stream for the first time
	ZegoStreamEventPlayFail

	// Start retrying playing stream
	ZegoStreamEventRetryPlayStart

	// Retry playing stream successfully
	ZegoStreamEventRetryPlaySuccess

	// Failed to retry playing stream
	ZegoStreamEventRetryPlayFail

	// End of playing stream
	ZegoStreamEventPlayEnd
)

// Playing stream status.
type ZegoPlayerState int

const (
	// The state of the flow is not played, and it is in this state before the stream is played. If the steady flow anomaly occurs during the playing process, such as AppID or Token are incorrect, it will enter this state.
	ZegoPlayerStateNoPlay ZegoPlayerState = iota

	// The state that the stream is being requested for playing. After the [startPlayingStream] function is successfully called, it will enter the state. The UI is usually displayed through this state. If the connection is interrupted due to poor network quality, the SDK will perform an internal retry and will return to the requesting state.
	ZegoPlayerStatePlayRequesting

	// The state that the stream is being playing, entering the state indicates that the stream has been successfully played, and the user can communicate normally.
	ZegoPlayerStatePlaying
)

// Player state.
type ZegoMediaPlayerState int

const (
	// Not playing
	ZegoMediaPlayerStateNoPlay ZegoMediaPlayerState = iota

	// Playing
	ZegoMediaPlayerStatePlaying

	// Pausing
	ZegoMediaPlayerStatePauing

	// End of play
	ZegoMediaPlayerStatePlayEnded
)

// Player network event.
type ZegoMediaPlayerNetworkEvent int

const (
	// Network resources are not playing well, and start trying to cache data
	ZegoMediaPlayerNetworkEventBufferBegin ZegoMediaPlayerNetworkEvent = iota

	// Network resources can be played smoothly
	ZegoMediaPlayerNetworkEventBufferEnded
)

// Media player first frame event type.
type ZegoMediaPlayerFirstFrameEvent int

const (
	// The first video frame is rendered event.
	ZegoMediaPlayerFirstFrameEventVideoRendered ZegoMediaPlayerFirstFrameEvent = iota

	// The first audio frame is rendered event.
	ZegoMediaPlayerFirstFrameEventAudioRendered
)

// Profile for create engine
//
// Profile for create engine
type ZegoEngineProfile struct {
	// Application ID issued by ZEGO for developers, please apply from the ZEGO Admin Console https://console.zegocloud.com The value ranges from 0 to 4294967295.
	AppID uint32

	// Application signature for each AppID, please apply from the ZEGO Admin Console. Application signature is a 64 character string. Each character has a range of '0' ~ '9', 'a' ~ 'z'. AppSign 2.17.0 and later allows null or no transmission. If the token is passed empty or not passed, the token must be entered in the [ZegoRoomConfig] parameter for authentication when the [loginRoom] interface is called to login to the room.
	AppSign string

	// The room scenario. the SDK will optimize the audio and video configuration for the specified scenario to achieve the best effect in this scenario. After specifying the scenario, you can call other APIs to adjusting the audio and video configuration. Differences between scenarios and how to choose a suitable scenario, please refer to https://docs.zegocloud.com/article/14940
	Scenario ZegoScenario
}

// Advanced engine configuration.
type ZegoEngineConfig struct {
	// @deprecated This property has been deprecated since version 2.3.0, please use the [setLogConfig] function instead.
	LogConfig *ZegoLogConfig

	// Other special function switches, if not set, no special function will be used by default. Please contact ZEGO technical support before use.
	AdvancedConfig map[string]string
}

// Log config.
//
// Description: This parameter is required when calling [setlogconfig] to customize log configuration.
// Use cases: This configuration is required when you need to customize the log storage path or the maximum log file size.
// Caution: None.
type ZegoLogConfig struct {
	// The storage path of the log file. Description: Used to customize the storage path of the log file. Use cases: This configuration is required when you need to customize the log storage path. Required: False. Default value: The default path of each platform is different, please refer to the official website document https://docs.zegocloud.com/faq/express_sdkLog. Caution: Developers need to ensure read and write permissions for files under this path.
	LogPath string

	// Maximum log file size(Bytes). Description: Used to customize the maximum log file size. Use cases: This configuration is required when you need to customize the upper limit of the log file size. Required: False. Default value: 5MB (5 * 1024 * 1024 Bytes). Value range: Minimum 1MB (1 * 1024 * 1024 Bytes), maximum 100M (100 * 1024 * 1024 Bytes), 0 means no need to write logs. Caution: The larger the upper limit of the log file size, the more log information it carries, but the log upload time will be longer.
	LogSize uint64

	// Log files count. Default is 3. Value range is [3, 20].
	LogCount uint32
}

func NewZegoLogConfig() ZegoLogConfig {
	return ZegoLogConfig{
		LogSize:  5 * 1024 * 1024,
		LogCount: 3,
	}
}

// User object.
//
// Configure user ID and username to identify users in the room.
// Note that the userID must be unique under the same appID, otherwise, there will be mutual kicks when logging in to the room.
// It is strongly recommended that userID corresponds to the user ID of the business APP, that is, a userID and a real user are fixed and unique, and should not be passed to the SDK in a random userID. Because the unique and fixed userID allows ZEGO technicians to quickly locate online problems.
type ZegoUser struct {
	// User ID, a utf8 string with a maximum length of 64 bytes or less.Privacy reminder: Please do not fill in sensitive user information in this field, including but not limited to mobile phone number, ID number, passport number, real name, etc.Caution: Only support numbers, English characters and '~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+', '=', '-', '`', ';', '’', ',', '.', '<', '>', '\'. Do not use '%' if you need to communicate with the Web SDK.
	UserID string

	// User Name, a utf8 string with a maximum length of 256 bytes or less.Please do not fill in sensitive user information in this field, including but not limited to mobile phone number, ID number, passport number, real name, etc.
	UserName string
}

// Advanced room configuration.
//
// Configure maximum number of users in the room and authentication token, etc.
type ZegoRoomConfig struct {
	// The maximum number of users in the room, Passing 0 means unlimited, the default is unlimited.
	MaxMemberCount uint32

	// Whether to enable the user in and out of the room callback notification [onRoomUserUpdate], the default is off. If developers need to use ZEGO Room user notifications, make sure that each user who login sets this flag to true
	IsUserStatusNotify bool

	// The token issued by the developer's business server is used to ensure security. For the generation rules, please refer to [Using Token Authentication](https://doc-zh.zego.im/article/10360), the default is an empty string, that is, no authentication. In versions 2.17.0 and above, if appSign is not passed in when calling the [createEngine] API to create an engine, or if appSign is empty, this parameter must be set for authentication when logging in to a room.
	Token string

	// The bitmask marker for capability negotiation, refer to enum [ZegoRoomCapabilityNegotiationTypesBitMask], when this param converted to binary, 0b01 that means 1 << 0 for enable the capability negotiation of all user in the room, 0x10 that means 1 << 1 for enable the capability negotiation of publisher in the room. The masks can be combined to allow different types of capability negotiation.
	CapabilityNegotiationTypes uint32

	// The type of the room, the default is 0.
	RoomType uint32
}

// Stream object.
//
// Identify an stream object
type ZegoStream struct {
	// User object instance.Please do not fill in sensitive user information in this field, including but not limited to mobile phone number, ID number, passport number, real name, etc.
	User ZegoUser

	// Stream ID, a string of up to 256 characters. Caution: You cannot include URL keywords, otherwise publishing stream and playing stream will fails. Only support numbers, English characters and '-', '_'.
	StreamID string

	// Stream extra info
	ExtraInfo string
}

// Advanced publisher configuration.
//
// Configure room id
type ZegoPublisherConfig struct {
	// The Room ID, It is not necessary to pass in single room mode, but the ID of the corresponding room must be passed in multi-room mode
	RoomID string

	// Whether to synchronize the network time when pushing streams. 1 is synchronized with 0 is not synchronized. And must be used with setStreamAlignmentProperty. It is used to align multiple streams at the mixed stream service or streaming end, such as the chorus scene of KTV.
	ForceSynchronousNetworkTime int

	// When pushing a flow, review the pattern of the flow. By default, no audit is performed. If you want to use this function, contact ZEGO technical support.
	StreamCensorshipMode ZegoStreamCensorshipMode

	// Inspect flag. If you want to use this function, contact ZEGO technical support.
	StreamCensorFlag int

	// Codec capability negotiation type. By default, no reference to the outcome of the capability negotiation. If you want to use this function, contact ZEGO technical support.
	CodecNegotiationType ZegoCapabilityNegotiationType

	// Stream title, a utf8 string with a maximum length of 255 bytes or less.
	StreamTitle string
}

// Audio configuration.
//
// Configure audio bitrate, audio channel, audio encoding for publishing stream
type ZegoAudioConfig struct {
	// Audio bitrate in kbps, default is 48 kbps. The settings before and after publishing stream can be effective
	Bitrate int

	// Audio channel, default is Mono. The setting only take effect before publishing stream
	Channel ZegoAudioChannel

	// codec ID, default is ZegoAudioCodecIDDefault. The setting only take effect before publishing stream
	CodecID ZegoAudioCodecID
}

// Custom audio configuration.
type ZegoCustomAudioConfig struct {
	// Audio capture source type
	SourceType ZegoAudioSourceType
}

// Parameter object for audio frame.
//
// Including the sampling rate and channel of the audio frame
type ZegoAudioFrameParam struct {
	// Sampling Rate
	SampleRate ZegoAudioSampleRate

	// Audio channel, default is Mono
	Channel ZegoAudioChannel
}

// Published stream quality information.
//
// Audio and video parameters and network quality, etc.
type ZegoPublishStreamQuality struct {
	// Video capture frame rate. The unit of frame rate is f/s
	VideoCaptureFPS float64

	// Video encoding frame rate. The unit of frame rate is f/s
	VideoEncodeFPS float64

	// Video transmission frame rate. The unit of frame rate is f/s
	VideoSendFPS float64

	// Video bit rate in kbps
	VideoKBPS float64

	// Audio capture frame rate. The unit of frame rate is f/s
	AudioCaptureFPS float64

	// Audio transmission frame rate. The unit of frame rate is f/s
	AudioSendFPS float64

	// Audio bit rate in kbps
	AudioKBPS float64

	// Local to server delay, in milliseconds
	Rtt int

	// Packet loss rate, in percentage, 0.0 ~ 1.0
	PacketLostRate float64

	// Published stream quality level
	Level ZegoStreamQualityLevel

	// Whether to enable hardware encoding
	IsHardwareEncode bool

	// Video codec ID (Available since 1.17.0)
	VideoCodecID ZegoVideoCodecID

	// Total number of bytes sent, including audio, video, SEI
	TotalSendBytes float64

	// Number of audio bytes sent
	AudioSendBytes float64

	// Number of video bytes sent
	VideoSendBytes float64
}

// Played stream quality information.
//
// Audio and video parameters and network quality, etc.
type ZegoPlayStreamQuality struct {
	// Video receiving frame rate. The unit of frame rate is f/s
	VideoRecvFPS float64

	// Video dejitter frame rate. The unit of frame rate is f/s (Available since 1.17.0)
	VideoDejitterFPS float64

	// Video decoding frame rate. The unit of frame rate is f/s
	VideoDecodeFPS float64

	// Video rendering frame rate. The unit of frame rate is f/s
	VideoRenderFPS float64

	// Video bit rate in kbps
	VideoKBPS float64

	// Video break rate, the unit is (number of breaks / every 10 seconds) (Available since 1.17.0)
	VideoBreakRate float64

	// Audio receiving frame rate. The unit of frame rate is f/s
	AudioRecvFPS float64

	// Audio dejitter frame rate. The unit of frame rate is f/s (Available since 1.17.0)
	AudioDejitterFPS float64

	// Audio decoding frame rate. The unit of frame rate is f/s
	AudioDecodeFPS float64

	// Audio rendering frame rate. The unit of frame rate is f/s
	AudioRenderFPS float64

	// Audio bit rate in kbps
	AudioKBPS float64

	// Audio break rate, the unit is (number of breaks / every 10 seconds) (Available since 1.17.0)
	AudioBreakRate float64

	// The audio quality of the playing stream determined by the audio MOS (Mean Opinion Score) measurement method, value range [-1, 5], where -1 means unknown, [0, 5] means valid score, the higher the score, the better the audio quality. For the subjective perception corresponding to the MOS value, please refer to https://docs.zegocloud.com/article/3720#4_4 (Available since 2.16.0)
	Mos float64

	// Server to local delay, in milliseconds
	Rtt int

	// Packet loss rate, in percentage, 0.0 ~ 1.0
	PacketLostRate float64

	// Delay from peer to peer, in milliseconds
	PeerToPeerDelay int

	// Packet loss rate from peer to peer, in percentage, 0.0 ~ 1.0
	PeerToPeerPacketLostRate float64

	// Published stream quality level
	Level ZegoStreamQualityLevel

	// Delay after the data is received by the local end, in milliseconds
	Delay int

	// The difference between the video timestamp and the audio timestamp, used to reflect the synchronization of audio and video, in milliseconds. This value is less than 0 means the number of milliseconds that the video leads the audio, greater than 0 means the number of milliseconds that the video lags the audio, and 0 means no difference. When the absolute value is less than 200, it can basically be regarded as synchronized audio and video, when the absolute value is greater than 200 for 10 consecutive seconds, it can be regarded as abnormal (Available since 1.19.0)
	AvTimestampDiff int

	// Whether to enable hardware decoding
	IsHardwareDecode bool

	// Video codec ID (Available since 1.17.0)
	VideoCodecID ZegoVideoCodecID

	// Total number of bytes received, including audio, video, SEI
	TotalRecvBytes float64

	// Number of audio bytes received
	AudioRecvBytes float64

	// Number of video bytes received
	VideoRecvBytes float64

	// Accumulated audio break count (Available since 2.9.0)
	AudioCumulativeBreakCount uint32

	// Accumulated audio break time, in milliseconds (Available since 2.9.0)
	AudioCumulativeBreakTime uint32

	// Accumulated audio break rate, in percentage, 0.0 ~ 100.0 (Available since 2.9.0)
	AudioCumulativeBreakRate float64

	// Accumulated audio decode time, in milliseconds (Available since 2.9.0)
	AudioCumulativeDecodeTime uint32

	// Accumulated video break count (Available since 2.9.0)
	VideoCumulativeBreakCount uint32

	// Accumulated video break time, in milliseconds (Available since 2.9.0)
	VideoCumulativeBreakTime uint32

	// Accumulated video break rate, in percentage, 0.0 ~ 1.0 (Available since 2.9.0)
	VideoCumulativeBreakRate float64

	// Accumulated video decode time, in milliseconds (Available since 2.9.0)
	VideoCumulativeDecodeTime uint32

	// Mute video (Available since 3.13.0)
	MuteVideo int

	// Mute audio (Available since 3.13.0)
	MuteAudio int
}

// SEI Callback info.
type ZegoMediaSideInfo struct {
	// Stream ID.
	StreamID string

	// SEI data
	SEIData []uint8

	// timestamp
	TimestampNs int64

	// SEI source module. Please contact ZEGO technical support.
	ModuleType int
}

// Room stream list.
//
// Room stream list.
type ZegoRoomStreamList struct {
	// Publish stream list
	PublishStreamList []ZegoStream

	// Play stream list
	PlayStreamList []ZegoStream
}

type IZegoEventHandler interface {
	// The callback for obtaining debugging error information.
	//
	// Available since: 1.1.0
	// Description: When the SDK functions are not used correctly, the callback prompts for detailed error information.
	// Trigger: Notify the developer when an exception occurs in the SDK.
	// Restrictions: None.
	// Caution: None.
	//
	// @param errorCode Error code, please refer to the error codes document https://docs.zegocloud.com/en/5548.html for details.
	// @param funcName Function name.
	// @param info Detailed error information.
	OnDebugError(errorCode int, funcName string, info string)

	// Notification of the room connection state changes.
	//
	// Available since: 1.1.0
	// Description: This callback is triggered when the connection status of the room changes, and the reason for the change is notified.For versions 2.18.0 and above, it is recommended to use the onRoomStateChanged callback instead of the onRoomStateUpdate callback to monitor room state changes.
	// Use cases: Developers can use this callback to determine the status of the current user in the room.
	// When to trigger:
	//  1. The developer will receive this notification when calling the [loginRoom], [logoutRoom], [switchRoom] functions.
	//  2. This notification may also be received when the network condition of the user's device changes (SDK will automatically log in to the room when disconnected, please refer to [Does ZEGO SDK support a fast reconnection for temporary disconnection] for details](https://docs.zegocloud.com/faq/reconnect?product=ExpressVideo&platform=all).
	// Restrictions: None.
	// Caution: If the connection is being requested for a long time, the general probability is that the user's network is unstable.
	// Related APIs: [loginRoom]、[logoutRoom]、[switchRoom]
	//
	// @param roomID Room ID, a string of up to 128 bytes in length.
	// @param state Changed room state.
	// @param errorCode Error code, For details, please refer to [Common Error Codes](https://docs.zegocloud.com/article/5548).
	// @param extendedData Extended Information with state updates. When the room login is successful, the key "room_session_id" can be used to obtain the unique RoomSessionID of each audio and video communication, which identifies the continuous communication from the first user in the room to the end of the audio and video communication. It can be used in scenarios such as call quality scoring and call problem diagnosis.
	OnRoomStateUpdate(roomID string, state ZegoRoomState, errorCode int, extendedData string)

	// The callback triggered when the number of streams published by the other users in the same room increases or decreases.
	//
	// Available since: 1.1.0
	// Description: When other users in the room start publishing stream or stop publishing stream, the streaming list in the room changes, and the developer will be notified through this callback.
	// Use cases: This callback is used to monitor stream addition or stream deletion notifications of other users in the room. Developers can use this callback to determine whether other users in the same room start or stop publishing stream, so as to achieve active playing stream [startPlayingStream] or take the initiative to stop the playing stream [stopPlayingStream], and use it to change the UI controls at the same time.
	// When to trigger:
	//   1. When the user logs in to the room for the first time, if there are other users publishing streams in the room, the SDK will trigger a callback notification with `updateType` being [ZegoUpdateTypeAdd], and `streamList` is an existing stream list.
	//   2. The user is already in the room. if another user adds a new push, the SDK will trigger a callback notification with `updateType` being [ZegoUpdateTypeAdd].
	//   3. The user is already in the room. If other users stop streaming, the SDK will trigger a callback notification with `updateType` being [ZegoUpdateTypeDelete].
	//   4. The user is already in the room. If other users leave the room, the SDK will trigger a callback notification with `updateType` being [ZegoUpdateTypeDelete].
	// Restrictions: None.
	//
	// @param roomID Room ID where the user is logged in, a string of up to 128 bytes in length.
	// @param updateType Update type (add/delete).
	// @param streamList Updated stream list.
	// @param extendedData Extended information with stream updates.When receiving a stream deletion notification, the developer can convert the string into a json object to get the stream_delete_reason field, which is an array of stream deletion reasons, and the stream_delete_reason[].code field may have the following values: 1 (the user actively stops publishing stream) ; 2 (user heartbeat timeout); 3 (user repeated login); 4 (user kicked out); 5 (user disconnected); 6 (removed by the server).
	OnRoomStreamUpdate(roomID string, updateType ZegoUpdateType, streamList []ZegoStream, extendedData string)

	// Notification of the room connection state changes, including specific reasons for state change.
	//
	// Available since: 2.18.0
	// Description: This callback is triggered when the connection status of the room changes, and the reason for the change is notified.For versions 2.18.0 and above, it is recommended to use the onRoomStateChanged callback instead of the onRoomStateUpdate callback to monitor room state changes.
	// Use cases: Developers can use this callback to determine the status of the current user in the room.
	// When to trigger: Users will receive this notification when they call room functions (refer to [Related APIs]). 2. This notification may also be received when the user device's network conditions change (SDK will automatically log in to the room again when the connection is disconnected, refer to https://doc-zh.zego.im/faq/reconnect ).
	// Restrictions: None.
	// Caution: If the connection is being requested for a long time, the general probability is that the user's network is unstable.
	// Related APIs: [loginRoom], [logoutRoom], [switchRoom]
	//
	// @param roomID Room ID, a string of up to 128 bytes in length.
	// @param reason Room state change reason.
	// @param errorCode Error code, please refer to the error codes document https://doc-en.zego.im/en/5548.html for details.
	// @param extendedData Extended Information with state updates. When the room login is successful, the key "room_session_id" can be used to obtain the unique RoomSessionID of each audio and video communication, which identifies the continuous communication from the first user in the room to the end of the audio and video communication. It can be used in scenarios such as call quality scoring and call problem diagnosis.
	OnRoomStateChanged(roomID string, reason ZegoRoomStateChangedReason, errorCode int, extendedData string)

	// Callback notification that room Token authentication is about to expire.
	//
	// Available since: 2.8.0
	// Description: The callback notification that the room Token authentication is about to expire, please use [renewToken] to update the room Token authentication.
	// Use cases: In order to prevent illegal entry into the room, it is necessary to perform authentication control on login room, push streaming, etc., to improve security.
	// When to call /Trigger: 30 seconds before the Token expires, the SDK will call [onRoomTokenWillExpire] to notify developer.
	// Restrictions: None.
	// Caution: The token contains important information such as the user's room permissions, publish stream permissions, and effective time, please refer to https://docs.zegocloud.com/article/11649.
	// Related APIs: When the developer receives this callback, he can use [renewToken] to update the token authentication information.
	//
	// @param roomID Room ID where the user is logged in, a string of up to 128 bytes in length.
	// @param remainTimeInSecond The remaining time before the token expires.
	OnRoomTokenWillExpire(roomID string, remainTimeInSecond int)

	// The callback triggered when the state of stream publishing changes.
	//
	// Available since: 1.1.0
	// Description: After calling the [startPublishingStream] successfully, the notification of the publish stream state change can be obtained through the callback function. You can roughly judge the user's uplink network status based on whether the state parameter is in [PUBLISH_REQUESTING].
	// Caution: The parameter [extendedData] is extended information with state updates. If you use ZEGO's CDN content distribution network, after the stream is successfully published, the keys of the content of this parameter are [flv_url_list], [rtmp_url_list], [hls_url_list], these correspond to the publishing stream URLs of the flv, rtmp, and hls protocols.
	// Related callbacks: After calling the [startPlayingStream] successfully, the notification of the play stream state change can be obtained through the callback function [onPlayerStateUpdate]. You can roughly judge the user's downlink network status based on whether the state parameter is in [PLAY_REQUESTING].
	//
	// @param streamID Stream ID.
	// @param state State of publishing stream.
	// @param errorCode The error code corresponding to the status change of the publish stream, please refer to the error codes document https://docs.zegocloud.com/en/5548.html for details.
	// @param extendedData Extended information with state updates, include playing stream CDN address.
	OnPublisherStateUpdate(streamID string, state ZegoPublisherState, errorCode int, extendedData string)

	// Callback for current stream publishing quality.
	//
	// Available since: 1.1.0
	// Description: After calling the [startPublishingStream] successfully, the callback will be received every 3 seconds default(If you need to change the time, please contact the instant technical support to configure). Through the callback, the collection frame rate, bit rate, RTT, packet loss rate and other quality data of the published audio and video stream can be obtained, and the health of the publish stream can be monitored in real time.You can monitor the health of the published audio and video streams in real time according to the quality parameters of the callback function, in order to show the uplink network status in real time on the device UI.
	// Caution: If you does not know how to use the parameters of this callback function, you can only pay attention to the [level] field of the [quality] parameter, which is a comprehensive value describing the uplink network calculated by SDK based on the quality parameters.
	// Related callbacks: After calling the [startPlayingStream] successfully, the callback [onPlayerQualityUpdate] will be received every 3 seconds. You can monitor the health of play streams in real time based on quality data such as frame rate, code rate, RTT, packet loss rate, etc.
	//
	// @param streamID Stream ID.
	// @param quality Publishing stream quality, including audio and video framerate, bitrate, RTT, etc.
	OnPublisherQualityUpdate(streamID string, quality ZegoPublishStreamQuality)

	// The callback triggered when publishing stream.
	//
	// Available since: 2.18.0
	// Description: After start publishing stream, this callback will return the current stream address, resource type and protocol-related information.
	// When to trigger: Publish and retry publish events.
	// Caution: None.
	//
	// @param eventID Publish stream event ID
	// @param streamID Stream ID.
	// @param extraInfo extra info. it is in JSON format. Included information includes "url" for address, "streamProtocol" for stream protocol, including rtmp, flv, avertp, hls, webrtc, etc. "netProtocol" for network protocol, including tcp, udp, quic, "resourceType" for resource type , including cdn, rtc, l3.
	OnPublisherStreamEvent(eventID ZegoStreamEvent, streamID string, extendedData string)

	// The callback triggered when the first audio frame is sent.
	//
	// Available since: 3.5.0
	// Description: After the [startPublishingStream] function is called successfully, this callback will be called when SDK received the first frame of audio data. Developers can use this callback to determine whether SDK has actually collected audio data. If the callback is not received, the audio capture device is occupied or abnormal.
	// Trigger: In the case of no startPublishingStream audio stream, the first startPublishingStream audio stream, it will receive this callback.
	// Related callbacks: After the [startPublishingStream] function is called successfully, determine if the SDK actually collected video data by the callback function [onPublisherCapturedVideoFirstFrame], determine if the SDK has rendered the first frame of video data collected by calling back [onPublisherRenderVideoFirstFrame].
	//
	// @param channel Publishing stream channel.If you only publish one audio stream, you can ignore this parameter.
	OnPublisherSendAudioFirstFrame(channel ZegoPublishChannel)

	// The callback triggered when the state of stream playing changes.
	//
	// Available since: 1.1.0
	// Description: After calling the [startPlayingStream] successfully, the notification of the playing stream state change can be obtained through the callback function. You can roughly judge the user's downlink network status based on whether the state parameter is in [PLAY_REQUESTING].
	// When to trigger:  After calling the [startPublishingStream], this callback is triggered when a playing stream's state changed.
	// Related callbacks: After calling the [startPublishingStream] successfully, the notification of the publish stream state change can be obtained through the callback function [onPublisherStateUpdate]. You can roughly judge the user's uplink network status based on whether the state parameter is in [PUBLISH_REQUESTING].
	//
	// @param streamID stream ID.
	// @param state State of playing stream.
	// @param errorCode The error code corresponding to the status change of the playing stream, please refer to the error codes document https://docs.zegocloud.com/en/5548.html for details.
	// @param extendedData Extended Information with state updates. As the standby, only an empty json table is currently returned.
	OnPlayerStateUpdate(streamID string, state ZegoPlayerState, errorCode int, extendedData string)

	// Callback for current stream playing quality.
	//
	// Available since: 1.1.0
	// Description: After calling the [startPlayingStream] successfully, the callback will be received every 3 seconds default(If you need to change the time, please contact the instant technical support to configure). Through the callback, the collection frame rate, bit rate, RTT, packet loss rate and other quality data can be obtained, and the health of the played audio and video streams can be monitored in real time.
	// Use cases: You can monitor the health of the played audio and video streams in real time according to the quality parameters of the callback function, in order to show the downlink network status on the device UI in real time.
	// Caution: If you does not know how to use the various parameters of the callback function, you can only focus on the level field of the quality parameter, which is a comprehensive value describing the downlink network calculated by SDK based on the quality parameters.
	// Related callbacks: After calling the [startPublishingStream] successfully, a callback [onPublisherQualityUpdate] will be received every 3 seconds. You can monitor the health of publish streams in real time based on quality data such as frame rate, code rate, RTT, packet loss rate, etc.
	//
	// @param streamID Stream ID.
	// @param quality Playing stream quality, including audio and video framerate, bitrate, RTT, etc.
	OnPlayerQualityUpdate(streamID string, quality ZegoPlayStreamQuality)

	// The callback triggered when Supplemental Enhancement Information is received synchronously.
	//
	// Available since: 3.9.0
	// Description: After the [startPlayingStream] function is called successfully, when the remote stream sends SEI (such as directly calling [sendSEI], audio mixing with SEI data, and sending custom video capture encoded data with SEI, etc.), the local end will receive this callback.
	// Trigger: After the [startPlayingStream] function is called successfully, when the remote stream sends SEI, the local end will receive this callback.
	// Caution: 1. Since the video encoder itself generates an SEI with a payload type of 5, or when a video file is used for publishing, such SEI may also exist in the video file. Therefore, if the developer needs to filter out this type of SEI, it can be before [createEngine] Call [ZegoEngineConfig.advancedConfig("unregister_sei_filter", "XXXXX")]. Among them, unregister_sei_filter is the key, and XXXXX is the uuid filter string to be set. 2. When [mutePlayStreamVideo] or [muteAllPlayStreamVideo] is called to set only the audio stream to be pulled, the SEI will not be received.
	//
	// @param info SEI Callback info.
	OnPlayerRecvMediaSideInfo(info ZegoMediaSideInfo)

	// The callback triggered when playing stream.
	//
	// Available since: 2.18.0
	// Description: After start playing stream, this callback will return the current stream address, resource type and protocol-related information.
	// When to trigger: Play and retry play events.
	// Caution: None.
	//
	// @param eventID Play stream event ID
	// @param streamID Stream ID.
	// @param extraInfo extra info. it is in JSON format. Included information includes "url" for address, "streamProtocol" for stream protocol, including rtmp, flv, avertp, hls, webrtc, etc. "netProtocol" for network protocol, including tcp, udp, quic, "resourceType" for resource type , including cdn, rtc, l3.
	OnPlayerStreamEvent(eventID ZegoStreamEvent, streamID string, extendedData string)

	// The callback triggered when the first audio frame is received.
	//
	// Available since: 1.1.0
	// Description: After the [startPlayingStream] function is called successfully, this callback will be called when SDK received the first frame of audio data.
	// Use cases: Developer can use this callback to count time consuming that take the first frame time or update the UI for playing stream.
	// Trigger: This callback is triggered when SDK receives the first frame of audio data from the network.
	// Related callbacks: After a successful call to [startPlayingStream], the callback function [onPlayerRecvVideoFirstFrame] determines whether the SDK has received the video data, and the callback [onPlayerRenderVideoFirstFrame] determines whether the SDK has rendered the first frame of the received video data.
	//
	// @param streamID Stream ID.
	OnPlayerRecvAudioFirstFrame(streamID string)
}

type IZegoAudioDataHandler interface {
	// The callback for obtaining the audio data of each stream.
	//
	// Available: Since 1.1.0
	// Description: This function will call back the data corresponding to each playing stream. Different from [onPlaybackAudioData], the latter is the mixed data of all playing streams. If developers need to process a piece of data separately, they can use this callback.
	// When to trigger: On the premise of calling [setAudioDataHandler] to set up listening for this callback, calling [startAudioDataObserver] to set the mask 0x08 that is 1 << 3, and this callback will be triggered when the SDK audio and video engine starts to play the stream.
	// Restrictions: None.
	// Caution: This callback is a high-frequency callback, please do not perform time-consuming operations in this callback.
	//
	// @param data Audio data in PCM format.
	// @param dataLength Length of the data.
	// @param param Parameters of the audio frame.
	// @param streamID Corresponding stream ID.
	OnPlayerAudioData(data []uint8, param ZegoAudioFrameParam, streamID string)
}

type IZegoApiCalledEventHandler interface {
	// Method execution result callback
	//
	// Available since: 2.3.0
	// Description: When the monitoring is turned on through [setApiCalledCallback], the results of the execution of all methods will be called back through this callback.
	// Trigger: When the developer calls the SDK method, the execution result of the method is called back.
	// Restrictions: None.
	// Caution: It is recommended to monitor and process this callback in the development and testing phases, and turn off the monitoring of this callback after going online.
	//
	// @param errorCode Error code, please refer to the error codes document https://docs.zegocloud.com/en/5548.html for details.
	// @param funcName Function name.
	// @param info Detailed error information.
	OnApiCalledResult(errorCode int, funcName string, info string)
}

type IZegoMediaPlayerEventHandler interface {
	// MediaPlayer playback status callback.
	//
	// Available since: 1.3.4
	// Description: MediaPlayer playback status callback.
	// Trigger: The callback triggered when the state of the media player changes.
	// Restrictions: None.
	//
	// @param mediaPlayer Callback player object.
	// @param state Media player status.
	// @param errorCode Error code, please refer to the error codes document https://docs.zegocloud.com/en/5548.html for details.
	OnMediaPlayerStateUpdate(mediaPlayer IZegoMediaPlayer, state ZegoMediaPlayerState, errorCode int)

	// The callback triggered when the network status of the media player changes.
	//
	// Available since: 1.3.4
	// Description: The callback triggered when the network status of the media player changes.
	// Trigger: When the media player is playing network resources, this callback will be triggered when the status change of the cached data.
	// Restrictions: The callback will only be triggered when the network resource is played.
	// Related APIs: [setNetWorkBufferThreshold].
	//
	// @param mediaPlayer Callback player object.
	// @param networkEvent Network status event.
	OnMediaPlayerNetworkEvent(mediaPlayer IZegoMediaPlayer, networkEvent ZegoMediaPlayerNetworkEvent)

	// The callback to report the current playback progress of the media player.
	//
	// Available since: 1.3.4
	// Description: The callback triggered when the network status of the media player changes. Set the callback interval by calling [setProgressInterval]. When the callback interval is set to 0, the callback is stopped. The default callback interval is 1 second.
	// Trigger: When the media player is playing network resources, this callback will be triggered when the status change of the cached data.
	// Restrictions: None.
	// Related APIs: [setProgressInterval].
	//
	// @param mediaPlayer Callback player object.
	// @param millisecond Progress in milliseconds.
	OnMediaPlayerPlayingProgress(mediaPlayer IZegoMediaPlayer, millisecond uint64)

	// The callback to report the current rendering progress of the media player.
	//
	// Available since: 3.8.0
	// Description: The callback to report the current rendering progress of the media player. Set the callback interval by calling [setProgressInterval]. When the callback interval is set to 0, the callback is stopped. The default callback interval is 1 second.
	// Trigger: This callback will be triggered when the media player starts playing resources.
	// Restrictions: None.
	// Related APIs: [setProgressInterval].
	//
	// @param mediaPlayer Callback player object.
	// @param millisecond Progress in milliseconds.
	OnMediaPlayerRenderingProgress(mediaPlayer IZegoMediaPlayer, millisecond uint64)

	// The callback triggered when the media player got media side info.
	//
	// Available since: 2.2.0
	// Description: The callback triggered when the media player got media side info.
	// Trigger: When the media player starts playing media files, the callback is triggered if the SEI is resolved to the media file.
	// Caution: The callback does not actually take effect until call [setEventHandler] to set.
	//
	// @param mediaPlayer Callback player object.
	// @param data SEI content.
	// @param dataLength SEI content length.
	OnMediaPlayerRecvSEI(mediaPlayer IZegoMediaPlayer, data []uint8)

	// The callback triggered when the media player plays the first frame.
	//
	// Available since: 3.5.0
	// Description: The callback triggered when the media player plays the first frame.
	// Trigger: This callback is generated when the media player starts playing.
	// Caution: The callback does not actually take effect until call [setEventHandler] to set.
	// Related APIs: You need to call the [setPlayerCanvas] interface to set the view for the media player in order to receive the video first frame event callback.
	//
	// @param mediaPlayer Callback player object.
	// @param event The first frame callback event type.
	OnMediaPlayerFirstFrameEvent(mediaPlayer IZegoMediaPlayer, event ZegoMediaPlayerFirstFrameEvent)
}

type IZegoMediaPlayerAudioHandler interface {
	// The callback triggered when the media player throws out audio frame data.
	//
	// Available since: 1.3.4
	// Description: The callback triggered when the media player throws out audio frame data.
	// Trigger: The callback is generated when the media player starts playing.
	// Caution: The callback does not actually take effect until call [setAudioHandler] to set.
	// Restrictions: When playing copyrighted music, this callback will be disabled by default. If necessary, please contact ZEGO technical support.
	//
	// @param mediaPlayer Callback player object.
	// @param data Raw data of audio frames.
	// @param dataLength Data length.
	// @param param audio frame flip mode.
	OnAudioFrame(mediaPlayer IZegoMediaPlayer, data []uint8, param ZegoAudioFrameParam)
}

// Callback for media player loads resources.
//
// @param errorCode Error code, please refer to the error codes document https://docs.zegocloud.com/en/5548.html for details.
type ZegoMediaPlayerLoadResourceCallback func(errorCode int)

// Callback for media player seek to playback progress.
//
// @param errorCode Error code, please refer to the error codes document https://docs.zegocloud.com/en/5548.html for details.
type ZegoMediaPlayerSeekToCallback func(errorCode int)

type IZegoMediaPlayer interface {
	// Set event notification callback handler of the media player.
	//
	// Available since: 2.1.0
	// Description: Listen to the event notification callback of the media player.
	// Use Cases: You can change the media player UI widget according to the related event callback.
	// When to call: After the [ZegoMediaPlayer] instance created.
	// Restrictions: None.
	// Caution: Calling this function will overwrite the callback set by the last call to this function.
	//
	// @param handler Event callback handler for media player
	SetEventHandler(handler IZegoMediaPlayerEventHandler)

	// Set audio data callback handler of the media player.
	//
	// Available since: 2.1.0
	// Description: By setting this callback, the audio data of the media resource file played by the media player can be called back.
	// When to call: After the [ZegoMediaPlayer] instance created.
	// Restrictions: None.
	// Caution: When you no longer need to get the audio data, please call this function again to clear the handler to stop the audio data callback.
	//
	// @param handler Audio data callback handler for media player
	SetAudioHandler(handler IZegoMediaPlayerAudioHandler)

	// Load local or network media resource.
	//
	// Available: since 1.3.4
	// Description: Load media resources.
	// Use case: Developers can load the absolute path to the local resource or the URL of the network resource incoming.
	// When to call: It can be called after the engine by [createEngine] has been initialized and the media player has been created by [createMediaPlayer].
	// Related APIs: Resources can be loaded through the [loadResourceWithPosition] or [loadResourceFromMediaData] function.
	// Caution: If the mediaplayer has already loaded resources or is in the process of playing, please first call the [stop] interface to halt the playback, and then proceed to call the interface to load the media resources; failure to do so will result in an unsuccessful load.
	//
	// @param path The absolute resource path or the URL of the network resource and cannot be nullptr or "". Android can set this path string with Uri.
	// @param callback Notification of resource loading results
	LoadResource(path string, callback ZegoMediaPlayerLoadResourceCallback)

	// Start playing.
	//
	// You need to load resources before playing
	Start()

	// Stop playing.
	Stop()

	// Pause playing.
	Pause()

	// Resume playing.
	Resume()

	// Set the specified playback progress.
	//
	// Unit is millisecond
	//
	// @param millisecond Point in time of specified playback progress
	// @param callback The result notification of set the specified playback progress
	SeekTo(millisecond uint64, callback ZegoMediaPlayerSeekToCallback)

	// Whether to repeat playback.
	//
	// @param enable repeat playback flag. The default is false.
	EnableRepeat(enable bool)

	// Whether to mix the player's sound into the stream being published.
	//
	// This interface will only mix the media player sound into the main channel
	//
	// @param enable Aux audio flag. The default is false.
	EnableAux(enable bool)

	// Set mediaplayer volume. Both the local play volume and the publish volume are set.
	//
	// @param volume The range is 0 ~ 200. The default is 60.
	SetVolume(volume int)

	// Set mediaplayer local playback volume.
	//
	// @param volume The range is 0 ~ 200. The default is 60.
	SetPlayVolume(volume int)

	// Set mediaplayer publish volume.
	//
	// @param volume The range is 0 ~ 200. The default is 60.
	SetPublishVolume(volume int)

	// Gets the current local playback volume of the mediaplayer, the range is 0 ~ 200, with the default value of 60.
	//
	// @return current volume
	GetPlayVolume() int

	// Gets the current publish volume of the mediaplayer, the range is 0 ~ 200, with the default value of 60.
	//
	// @return current volume
	GetPublishVolume() int

	// Get media player index.
	//
	// Description: Get media player index.
	// When to call: It can be called after [createMediaPlayer].
	// Restrictions: None.
	//
	// @return Media player index.
	GetIndex() int
}

// Login room result callback.
//
// @param errorCode Error code, please refer to the error codes document https://docs.zegocloud.com/en/5548.html for details.
// @param extendedData Extended Information
type ZegoRoomLoginCallback func(errorCode int, extendedData string)

// Logout room result callback.
//
// @param errorCode Error code, please refer to the error codes document https://docs.zegocloud.com/en/5548.html for details.
// @param extendedData Extended Information
type ZegoRoomLogoutCallback func(errorCode int, extendedData string)

// Callback for sending broadcast messages.
//
// @param errorCode Error code, please refer to the error codes document https://docs.zegocloud.com/en/5548.html for details.
// @param messageID ID of this message
type ZegoIMSendBroadcastMessageCallback func(errorCode int, messageID uint64)

type IZegoExpressEngine interface {
	// Enable the debug assistant. Note, do not enable this feature in the online version! Use only during development phase!
	//
	// Available since: 2.17.0
	// Description: After enabled, the SDK will print logs to the console, and will pop-up an alert (toast) UI message when there is a problem with calling other SDK functions.
	// Default value: This function is disabled by default.
	// When to call: This function can be called right after [createEngine].
	// Platform differences: The pop-up alert function only supports Android / iOS / macOS / Windows, and the console log function supports all platforms.
	// Caution: Be sure to confirm that this feature is turned off before the app is released to avoid pop-up UI alert when an error occurs in your release version's app. It is recommended to associate the [enable] parameter of this function with the DEBUG variable of the app, that is, only enable the debug assistant in the DEBUG environment.
	// Restrictions: None.
	//
	// @param enable Whether to enable the debug assistant.
	EnableDebugAssistant(enable bool)

	// Logs in to a room with advanced room configurations. You must log in to a room before publishing or playing streams.
	//
	// Available since: 1.1.0
	// Description: If the room does not exist, [loginRoom] creates and logs in the room. SDK uses the 'room' to organize users. After users log in to a room, they can use interface such as push stream [startPublishingStream], pull stream [startPlayingStream], send and receive broadcast messages [sendBroadcastMessage], etc. To prevent the app from being impersonated by a malicious user, you can add authentication before logging in to the room, that is, the [token] parameter in the ZegoRoomConfig object passed in by the [config] parameter.
	// Use cases: In the same room, users can conduct live broadcast, audio and video calls, etc.
	// When to call /Trigger: This interface is called after [createEngine] initializes the SDK.
	// Restrictions: For restrictions on the use of this function, please refer to https://docs.zegocloud.com/article/7611 or contact ZEGO technical support.
	// Caution:
	//   1. Apps that use different appIDs cannot intercommunication with each other.
	//   2. SDK supports startPlayingStream audio and video streams from different rooms under the same appID, that is, startPlayingStream audio and video streams across rooms. Since ZegoExpressEngine's room related callback notifications are based on the same room, when developers want to startPlayingStream streams across rooms, developers need to maintain related messages and signaling notifications by themselves.
	//   3. It is strongly recommended that userID corresponds to the user ID of the business APP, that is, a userID and a real user are fixed and unique, and should not be passed to the SDK in a random userID. Because the unique and fixed userID allows ZEGO technicians to quickly locate online problems.
	//   4. After the first login failure due to network reasons or the room is disconnected, the default time of SDK reconnection is 20min.
	// Privacy reminder: Please do not fill in sensitive user information in this interface, including but not limited to mobile phone number, ID number, passport number, real name, etc.
	// Related callbacks:
	//   1. When the user starts to log in to the room, the room is successfully logged in, or the room fails to log in, the [onRoomStateChanged] (Not supported before 2.18.0, please use [onRoomStateUpdate]) callback will be triggered to notify the developer of the status of the current user connected to the room.
	//   2. Different users who log in to the same room can get room related notifications in the same room (eg [onRoomUserUpdate], [onRoomStreamUpdate], etc.), and users in one room cannot receive room signaling notifications in another room.
	//   3. If the network is temporarily interrupted due to network quality reasons, the SDK will automatically reconnect internally. You can get the current connection status of the local room by listening to the [onRoomStateChanged] (Not supported before 2.18.0, please use [onRoomStateUpdate]) callback method, and other users in the same room will receive [onRoomUserUpdate] callback notification.
	//   4. Messages sent in one room (e.g. [setStreamExtraInfo], [sendBroadcastMessage], [sendBarrageMessage], [sendCustomCommand], etc.) cannot be received callback ((eg [onRoomStreamExtraInfoUpdate], [onIMRecvBroadcastMessage], [onIMRecvBarrageMessage], [onIMRecvCustomCommand], etc) in other rooms. Currently, SDK does not provide the ability to send messages across rooms. Developers can integrate the SDK of third-party IM to achieve.
	// Related APIs:
	//   1. Users can call [logoutRoom] to log out. In the case that a user has successfully logged in and has not logged out, if the login interface is called again, the console will report an error and print the error code 1002001.
	//   2. SDK supports multi-room login, please call [setRoomMode] function to select multi-room mode before engine initialization, and then call [loginRoom] to log in to multi-room.
	//   3. Calling [destroyEngine] will also automatically log out.
	//
	// @param roomID Room ID, a string of less 128 bytes in length.
	//   Caution:
	//   1. room ID is defined by yourself.
	//   2. Only support numbers, English characters and '~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+', '=', '-', '`', ';', '’', ',', '.', '<', '>', '\'.
	//   3. If you need to communicate with the Web SDK, please do not use '%'.
	// @param user User object instance, configure userID, userName. Note that the userID needs to be globally unique with the same appID, otherwise the user who logs in later will kick out the user who logged in first.
	// @param config Advanced room configuration.
	// @param callback The callback of this login result(2.18.0 and above version support), if you need detailed room status, please pay attention to the [onRoomStateChanged] callback. Required: No. Default value: nullptr.Caution: If the connection is retried multiple times due to network problems, the retry status will not be thrown by this callback.
	LoginRoom(roomID string, user ZegoUser, config *ZegoRoomConfig, callback ZegoRoomLoginCallback)

	// Exit the room of the specified room ID with callback.
	//
	// Available since: 1.1.0
	// Description: This API will log out the room named roomID.
	// Use cases: In the same room, users can conduct live broadcast, audio and video calls, etc.
	// When to call /Trigger: After successfully logging in to the room, if the room is no longer used, the user can call the function [logoutRoom].
	// Restrictions: None.
	// Caution: 1. Exiting the room will stop all publishing and playing streams for user, and inner audio and video engine will stop, and then SDK will auto stop local preview UI. If you want to keep the preview ability when switching rooms, please use the [switchRoom] method. 2. If the user logs out to the room, but the incoming 'roomID' is different from the logged-in room name, SDK will return failure.
	// Related callbacks: After calling this function, you will receive [onRoomStateChanged] (Not supported before 2.18.0, please use [onRoomStateUpdate]) callback notification successfully exits the room, while other users in the same room will receive the [onRoomUserUpdate] callback notification(On the premise of enabling isUserStatusNotify configuration).
	// Related APIs: Users can use [loginRoom], [switchRoom] functions to log in or switch rooms.
	//
	// @param roomID Room ID, a string of up to 128 bytes in length.
	//   Caution:
	//   1. Only support numbers, English characters and '~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+', '=', '-', '`', ';', '’', ',', '.', '<', '>', '\'.
	//   2. If you need to communicate with the Web SDK, please do not use '%'.
	// @param callback The callback of this logout room result, if you need detailed room status, please pay attention to the [onRoomStateChanged] callback. Required: No. Default value: nullptr.
	LogoutRoom(roomID string, callback ZegoRoomLogoutCallback)

	// Renew token.
	//
	// Available since: 2.8.0
	// Description: After the developer receives [onRoomTokenWillExpire], they can use this API to update the token to ensure that the subsequent RTC functions are normal.
	// Use cases: Used when the token is about to expire.
	// When to call /Trigger: After the developer receives [onRoomTokenWillExpire].
	// Restrictions: None.
	// Caution: The token contains important information such as the user's room permissions, publish stream permissions, and effective time, please refer to https://docs.zegocloud.com/article/11649.
	// Related callbacks: None.
	// Related APIs: None.
	//
	// @param roomID Room ID.
	// @param token The token that needs to be renew.
	RenewToken(roomID string, token string)

	// Sends a Broadcast Message.
	//
	// Available since: 1.2.1
	// Description: Send a broadcast message to the room, users who have entered the same room can receive the message, and the message is reliable.
	// Use cases: Generally used in the live room.
	// When to call: After calling [loginRoom] to log in to the room.
	// Restrictions: The frequency of sending broadcast messages in the same room cannot be higher than 10 messages/s. The maximum QPS for a single user calling this interface from the client side is 2. For restrictions on the use of this function, please contact ZEGO technical support.
	// Related callbacks: The room broadcast message can be received through [onIMRecvBroadcastMessage].
	// Related APIs: Barrage messages can be sent through the [sendBarrageMessage] function, and custom command can be sent through the [sendCustomCommand] function.
	//
	// @param roomID Room ID, a string of less 128 bytes in length.
	//   Caution:
	//   1. room ID is defined by yourself.
	//   2. Only support numbers, English characters and '~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+', '=', '-', '`', ';', '’', ',', '.', '<', '>', '\'.
	//   3. If you need to communicate with the Web SDK, please do not use '%'.
	// @param message The content of the message. Required: Yes. Value range: The length does not exceed 1024 bytes.
	// @param callback Send a notification of the result of a broadcast message. Required: No. Caution: Passing [nullptr] means not receiving callback notifications.
	SendBroadcastMessage(roomID string, message string, callback ZegoIMSendBroadcastMessageCallback)

	// Get room stream list.
	//
	// Available since: 3.12.0
	// Description: Get room stream list.
	// Use cases: Get room stream list.
	// When to call /Trigger: After logging in the room successful.
	// Caution: This interface is to get a real-time internal stream list, which may be inaccurate when the room is disconnected from the service. Do not call this interface with high frequency.
	// Related APIs: None.
	//
	// @param roomID Room ID.
	// @param streamListType Get type
	// @return return stream list
	GetRoomStreamList(roomID string, streamListType ZegoRoomStreamListType) ZegoRoomStreamList

	// Starts publishing a stream. Support multi-room mode.
	//
	// Available since: 1.1.0
	// Description: Users push their local audio and video streams to the ZEGO RTC server or CDN, and other users in the same room can pull the audio and video streams to watch through the `streamID` or CDN pull stream address.
	// Use cases: It can be used to publish streams in real-time connecting wheat, live broadcast and other scenarios.
	// When to call: After [loginRoom].
	// Restrictions: None.
	// Caution:
	//   1. Before start to publish the stream, the user can choose to call [setVideoConfig] to set the relevant video parameters, and call [startPreview] to preview the video.
	//   2. Other users in the same room can get the streamID by monitoring the [onRoomStreamUpdate] event callback after the local user publishing stream successfully.
	//   3. In the case of poor network quality, user publish may be interrupted, and the SDK will attempt to reconnect. You can learn about the current state and error information of the stream published by monitoring the [onPublisherStateUpdate] event.
	//   4. To call [SetRoomMode] function to select multiple rooms, the room ID must be specified explicitly.
	//
	// @param streamID Stream ID, a string of up to 256 characters.
	//   Caution:
	//   1. Stream ID is defined by you.
	//   2. needs to be globally unique within the entire AppID. If in the same AppID, different users publish each stream and the stream ID is the same, which will cause the user to publish the stream failure. You cannot include URL keywords, otherwise publishing stream and playing stream will fails.
	//   3. Only support numbers, English characters and '-', '_'.
	// @param config Advanced publish configuration.
	// @param channel Publish stream channel.
	StartPublishingStream(streamID string, config ZegoPublisherConfig, channel ZegoPublishChannel)

	// Stops publishing a stream (for the specified channel).
	//
	// Available since: 1.1.0
	// Description: The user stops sending local audio and video streams, and other users in the room will receive a stream deletion notification.
	// Use cases: It can be used to stop publish streams in real-time connecting wheat, live broadcast and other scenarios.
	// When to call: After [startPublishingStream].
	// Restrictions: None.
	// Caution:
	//   1. After stopping the streaming, other users in the same room can receive the delete notification of the stream by listening to the [onRoomStreamUpdate] callback.
	//   2. If the user has initiated publish flow, this function must be called to stop the publish of the current stream before publishing the new stream (new streamID), otherwise the new stream publish will return a failure.
	//   3. After stopping streaming, the developer should stop the local preview based on whether the business situation requires it.
	//
	// @param channel Publish stream channel.
	StopPublishingStream(channel ZegoPublishChannel)

	// Sets up the audio configurations for the specified publish channel.
	//
	// Available since: 1.3.4
	// Description: You can set the combined value of the audio codec, bit rate, and audio channel through this function. If the preset value cannot meet the developer's scenario, the developer can set the parameters according to the business requirements.
	// Default value: The default audio config refers to the default value of [ZegoAudioConfig].
	// When to call: After the engine is created [createEngine], and before publishing [startPublishingStream].
	// Restrictions: None.
	// Related APIs: [getAudioConfig].
	//
	// @param config Audio config.
	// @param channel Publish stream channel.
	SetAudioConfig(config ZegoAudioConfig, channel ZegoPublishChannel)

	// Whether to enable acoustic echo cancellation (AEC).
	//
	// Available since: 1.1.0
	// Description: Turning on echo cancellation, the SDK filters the collected audio data to reduce the echo component in the audio.
	// Use case: When you need to reduce the echo to improve the call quality and user experience, you can turn on this feature.
	// When to call: It needs to be called after [createEngine].
	// Caution: The AEC function only supports the processing of sounds playbacked through the SDK, such as sounds played by the playing stream, media player, audio effect player, etc. Before this function is called, the SDK automatically determines whether to use AEC. Once this function is called, the SDK does not automatically determine whether to use AEC.
	// Restrictions: None.
	// Related APIs: Developers can use [enableHeadphoneAEC] to set whether to enable AEC when using headphones, and use [setAECMode] to set the echo cancellation mode.
	//
	// @param enable Whether to enable echo cancellation, true: enable, false: disable
	EnableAEC(enable bool)

	// Enables the custom audio I/O function (for the specified channel), support PCM, AAC format data.
	//
	// Available since: 1.10.0
	// Description: Enable custom audio IO function, support PCM, AAC format data.
	// Use cases: If the developer wants to implement special functions (such as voice change, bel canto, etc.) through custom processing after the audio data is collected or before the remote audio data is drawn for rendering.
	// When to call: It needs to be called before [startPublishingStream], [startPlayingStream], [startPreview], [createMediaPlayer], [createAudioEffectPlayer] and [createRealTimeSequentialDataManager] to be effective.
	// Restrictions: None.
	//
	// @param enable Whether to enable custom audio IO, default is false.
	// @param config Custom audio IO config.
	// @param channel Specify the publish channel to enable custom audio IO.
	EnableCustomAudioIO(enable bool, config *ZegoCustomAudioConfig, channel ZegoPublishChannel)

	// Sends Supplemental Enhancement Information to the specified publish channel.
	//
	// Available since: 1.1.0
	// Description: While pushing the stream to transmit the audio and video stream data, the stream media enhancement supplementary information is sent to synchronize some other additional information.
	// Use cases: Generally used in scenes such as synchronizing music lyrics or precise video layout, you can choose to send SEI.
	// When to call: After starting to push the stream [startPublishingStream].
	// Restrictions: Do not exceed 30 times per second, and the SEI data length is limited to 4096 bytes.
	// Caution: 1. Due to network issues, frame loss may occur, which means SEI information may also be lost. To address this situation, it is advisable to send it multiple times within a limited frequency. 2. Even if the [enableCamera] interface is called to turn off the camera or [mutePublishStreamVideo] is used to stop sending video data, SEI can still be successfully sent; as long as the playback side does not call the [mutePlayStreamVideo] interface to stop pulling audio data, SEI can still be received normally. 3. If the SDK does not support the video module but does support the SEI functional module, SEI information can still be sent normally.
	// Related APIs: After the pusher sends the SEI, the puller can obtain the SEI content by monitoring the callback of [onPlayerRecvSEI].
	//
	// @param data SEI data.
	// @param dataLength SEI data length.
	// @param channel Publish stream channel.
	SendSEI(data []uint8, channel ZegoPublishChannel)

	// Sets up the event callback handler for receiving audio data.
	//
	// Available since: 1.1.0
	// Description: This function can be called to receive audio data thrown by SDK bypass.
	// Use cases: When the developer needs to obtain the audio data of the remote user or the data collected by the local microphone for other purposes (such as pure audio recording, pure audio third-party monitoring, pure audio real-time analysis).
	// When to call: After creating the engine.
	// Restrictions: The set callback needs to be effective after calling [startAudioDataObserver] and is in the push or pull state.
	// Related APIs: Turn on the audio data monitoring call [startAudioDataObserver], turn off the audio data monitoring call [stopAudioDataObserver].
	// Caution: The format of the thrown audio data is pcm. The SDK still controls the collection and playback of the sound device. This function is to copy a copy of the data collected or played inside the SDK for use by the developer.
	//
	// @param handler Audio data handler for receive audio data.
	SetAudioDataHandler(handler IZegoAudioDataHandler)

	// Enable audio data observering.
	//
	// Available since: 1.1.0
	// Description: When custom audio processing is turned on, use this function to turn on audio data callback monitoring.
	// Use cases: When develop need to monitor the original audio.
	// When to call: After creating the engine.
	// Restrictions: Audio monitoring is triggered only after this function has been called and the callback has been set by calling [setAudioDataHandler]. If you want to enable the [onPlayerAudioData] callback, you must also be in the pull stream, and the incoming sampling rate of the [startAudioDataObserver] function is not supported at 8000Hz, 22050Hz, and 24000Hz.
	// Caution: This api will start the media engine and occupy the microphone device.
	//
	// @param observerBitMask The callback function bitmask marker for receive audio data, refer to enum [ZegoAudioDataCallbackBitMask], when this param converted to binary, 0b01 that means 1 << 0 for triggering [onCapturedAudioData], 0x10 that means 1 << 1 for triggering [onPlaybackAudioData], 0x100 that means 1 << 2 for triggering [onMixedAudioData], 0x1000 that means 1 << 3 for triggering [onPlayerAudioData]. The masks can be combined to allow different callbacks to be triggered simultaneously.
	// @param param param of audio frame.
	StartAudioDataObserver(observerBitMask uint32, param ZegoAudioFrameParam)

	// Disable audio data observering.
	//
	// Available since: 1.1.0
	// Description: Disable audio data observering.
	// Use cases: When develop need to monitor the original audio.
	// When to call: After calling [startAudioDataObserver] to start audio data monitoring.
	StopAudioDataObserver()

	// Starts playing a stream from ZEGO RTC server, without Canvas, it is more suitable for pure audio stream.
	//
	// Available since: 1.1.0
	// Description: Play audio streams from the ZEGO RTC server or CDN.
	// Use cases: In the real-time scenario, developers can listen to the [onRoomStreamUpdate] event callback to obtain the new stream information in the room where they are located, and call this interface to pass in streamID for play streams.
	// When to call: After [loginRoom].
	// Restrictions: None.
	// Caution: 1. After the first play stream failure due to network reasons or the play stream is interrupted, the default time for SDK reconnection is 20min. 2. In the case of poor network quality, user play may be interrupted, the SDK will try to reconnect, and the current play status and error information can be obtained by listening to the [onPlayerStateUpdate] event. please refer to https://docs.zegocloud.com/faq/reconnect. 3. Playing the stream ID that does not exist, the SDK continues to try to play after calling this function. After the stream ID is successfully published, the audio stream can be actually played.
	//
	// @param streamID Stream ID, a string of up to 256 characters.
	//   Caution:
	//   Only support numbers, English characters and '-', '_'.
	StartPlayingStream(streamID string)

	// Stops playing a stream.
	//
	// Available since: 1.1.0
	// Description: Play audio and video streams from the ZEGO RTC server.
	// Use cases: In the real-time scenario, developers can listen to the [onRoomStreamUpdate] event callback to obtain the delete stream information in the room where they are located, and call this interface to pass in streamID for stop play streams.
	// When to call: After [loginRoom].
	// Restrictions: None.
	// Caution: 1. When stopped, the attributes set for this stream previously, such as [setPlayVolume], [mutePlayStreamAudio], [mutePlayStreamVideo], etc., will be invalid and need to be reset when playing the the stream next time.
	//  2. After stopping pulling, the iOS platform view will clear the last frame by default and keep the background color of the view. The Android platform view remains at the last frame by default. If you need to clear the last frame, please contact ZEGO technical support.
	//
	// @param streamID Stream ID.
	StopPlayingStream(streamID string)

	// Sends PCM audio data produced by custom audio capture to the SDK (for the specified channel).
	//
	// Available since: 1.10.0
	// Description: Sends the captured audio PCM data to the SDK.
	// Use cases: 1.The customer needs to obtain input after acquisition from the existing audio stream, audio file, or customized acquisition system, and hand it over to the SDK for transmission. 2.Customers have their own requirements for special sound processing for PCM input sources. After the sound processing, the input will be sent to the SDK for transmission.
	// When to call: After [enableCustomAudioIO] and publishing stream successfully.
	// Restrictions: None.
	// Related APIs: Enable the custom audio IO function [enableCustomAudioIO], and start the push stream [startPublishingStream].
	//
	// @param data PCM buffer data.
	// @param dataLength The total length of the buffer data.
	// @param param The param of this PCM audio frame.
	// @param channel Publish channel for capturing audio frames.
	SendCustomAudioCapturePCMData(data []uint8, param ZegoAudioFrameParam, channel ZegoPublishChannel)

	// Fetches PCM audio data of the remote stream from the SDK for custom audio rendering.
	//
	// Available since: 1.10.0
	// Description: Fetches PCM audio data of the remote stream from the SDK for custom audio rendering, it is recommended to use the system framework to periodically invoke this function to drive audio data rendering.
	// Use cases: When developers have their own rendering requirements, such as special applications or processing and rendering of the original PCM data that are pulled, it is recommended to use the custom audio rendering function of the SDK.
	// When to call: After [enableCustomAudioIO] and playing stream successfully.
	// Restrictions: None.
	// Related APIs: Enable the custom audio IO function [enableCustomAudioIO], and start the play stream [startPlayingStream].
	//
	// @param data A block of memory for storing audio PCM data that requires user to manage the memory block's lifecycle, the SDK will copy the audio frame rendering data to this memory block.
	// @param dataLength The length of the audio data to be fetch this time (dataLength = duration * sampleRate * channels * 2(16 bit depth i.e. 2 Btye)).
	// @param param Specify the parameters of the fetched audio frame. sampleRate in ZegoAudioFrameParam must assignment
	FetchCustomAudioRenderPCMData(data []uint8, param ZegoAudioFrameParam)

	// Creates a media player instance.
	//
	// Available since: 2.1.0
	// Description: Creates a media player instance.
	// Use case: It is often used to play media resource scenes, For example, play video files, push the video of media resources in combination with custom video acquisition, and the remote end can pull the stream for viewing.
	// When to call: It can be called after the SDK by [createEngine] has been initialized.
	// Restrictions: Currently, a maximum of 4 instances can be created, after which it will return nullptr.
	// Caution: The more instances of a media player, the greater the performance overhead on the device.
	// Related APIs: User can call [destroyMediaPlayer] function to destroy a media player instance.
	//
	// @return Media player instance, nullptr will be returned when the maximum number is exceeded.
	CreateMediaPlayer() IZegoMediaPlayer

	// Destroys a media player instance.
	//
	// Available since: 2.1.0
	// Description: Destroys a media player instance.
	// Related APIs: User can call [createMediaPlayer] function to create a media player instance.
	//
	// @param mediaPlayer The media player instance object to be destroyed.
	DestroyMediaPlayer(mediaPlayer IZegoMediaPlayer)
}

// Create ZegoExpressEngine singleton object and initialize SDK.
//
// Available since: 2.14.0
// Description: Create ZegoExpressEngine singleton object and initialize SDK.
// When to call: The engine needs to be created before calling other functions.
// Restrictions: None.
// Caution: The SDK only supports the creation of one instance of ZegoExpressEngine. If you need call [createEngine] multiple times, you need call [destroyEngine] before you call the next [createEngine]. Otherwise it will return the instance which created by [createEngine] you called last time.
//
// @param profile The basic configuration information is used to create the engine.
// @param eventHandler Event notification callback. [nullptr] means not receiving any callback notifications.It can also be managed later via [setEventHandler]. If [createEngine] is called repeatedly and the [destroyEngine] function is not called to destroy the engine before the second call, the eventHandler will not be updated.
// @return engine singleton instance.
func CreateEngine(profile ZegoEngineProfile, eventHandler IZegoEventHandler) IZegoExpressEngine {
	return createEngine(profile, eventHandler)
}

// Callback for asynchronous destruction completion.
//
// In general, developers do not need to listen to this callback.
type ZegoDestroyCompletionCallback func()

// Destroy the ZegoExpressEngine singleton object and deinitialize the SDK.
//
// Available since: 1.1.0
// Description: Destroy the ZegoExpressEngine singleton object and deinitialize the SDK.
// When to call: When the SDK is no longer used, the resources used by the SDK can be released through this interface
// Restrictions: None.
// Caution: After using [createEngine] to create a singleton, if the singleton object has not been created or has been destroyed, you will not receive related callbacks when calling this function.
//
// @param engine engine instance that created by createEngine method.
// @param callback Notification callback for destroy engine completion. Developers can listen to this callback to ensure that device hardware resources are released. If the developer only uses SDK to implement audio and video functions, this parameter can be passed [nullptr].
func DestroyEngine(engine IZegoExpressEngine, callback ZegoDestroyCompletionCallback) {
	destroyEngine(engine, callback)
}

// Returns the singleton instance of ZegoExpressEngine.
//
// Available since: 1.1.0
// Description: If the engine has not been created or has been destroyed, returns [nullptr].
// When to call: After creating the engine, before destroying the engine.
// Restrictions: None.
//
// @return Engine singleton instance
func GetEngine() IZegoExpressEngine {
	return getEngine()
}

// Set advanced engine configuration.
//
// Available since: 1.1.0
// Description: Used to enable advanced functions.
// When to call: Different configurations have different call timing requirements. For details, please consult ZEGO technical support.
// Restrictions: None.
//
// @param config Advanced engine configuration
func SetEngineConfig(config ZegoEngineConfig) {
	setEngineConfig(config)
}

// Set log configuration.
//
// Available since: 2.3.0
// Description: If you need to customize the log file size and path, please call this function to complete the configuration.
// When to call: It must be set before calling [createEngine] to take effect. If it is set after [createEngine], it will take effect at the next [createEngine] after [destroyEngine].
// Restrictions: None.
// Caution: Once this interface is called, the method of setting log size and path via [setEngineConfig] will be invalid.Therefore, it is not recommended to use [setEngineConfig] to set the log size and path.
//
// @param config log configuration.
func SetLogConfig(config ZegoLogConfig) {
	setLogConfig(config)
}

// Set room mode.
//
// Available since: 2.9.0
// Description: If you need to use the multi-room feature, please call this function to complete the configuration.
// When to call: Must be set before calling [createEngine] to take effect, otherwise it will fail.
// Restrictions: If you need to use the multi-room feature, please contact the instant technical support to configure the server support.
// Caution: None.
//
// @param mode Room mode. Description: Used to set the room mode. Use cases: If you need to enter multiple rooms at the same time for publish-play stream, please turn on the multi-room mode through this interface. Required: True. Default value: ZEGO_ROOM_MODE_SINGLE_ROOM.
func SetRoomMode(mode ZegoRoomMode) {
	setRoomMode(mode)
}

// Gets the SDK's version number.
//
// Available since: 1.1.0
// Description: If you encounter an abnormality during the running of the SDK, you can submit the problem, log and other information to the ZEGO technical staff to locate and troubleshoot. Developers can also collect current SDK version information through this API, which is convenient for App operation statistics and related issues.
// When to call: Any time.
// Restrictions: None.
// Caution: None.
//
// @return SDK version.
func GetVersion() string {
	return getVersion()
}

// Set method execution result callback.
//
// Available since: 2.3.0
// Description: Set the setting of the execution result of the calling method. After setting, you can get the detailed information of the result of each execution of the ZEGO SDK method.
// When to call: Any time.
// Restrictions: None.
// Caution: It is recommended that developers call this interface only when they need to obtain the call results of each interface. For example, when troubleshooting and tracing problems. Developers generally do not need to pay attention to this interface.
//
// @param callback Method execution result callback.
func SetApiCalledCallback(callback IZegoApiCalledEventHandler) {
	setApiCalledCallback(callback)
}

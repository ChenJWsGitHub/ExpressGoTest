package zegoexpress

/*
#cgo CFLAGS: -I${SRCDIR}/lib/include
#cgo LDFLAGS: -L${SRCDIR}/lib -lZegoExpressEngine -Wl,-rpath=${SRCDIR}/lib
#include <stddef.h>
#include <stdlib.h>
#include <string.h>
#include "zego-express-errcode.h"
#include "zego-express-engine.h"
#include "zego-express-room.h"
#include "zego-express-publisher.h"
#include "zego-express-player.h"
#include "zego-express-im.h"
#include "zego-express-preprocess.h"
#include "zego-express-custom-audio-io.h"
#include "zego-express-device.h"
#include "zego-express-mediaplayer.h"

// 声明由Go实现的函数
extern void GoOnApiCalledResult(int, char*, char*);
extern void GoLoginResultCallback(int, char*, int);
extern void GoLogoutResultCallback(int, char*, int);
extern void GoOnIMSendBroadcastMessageResult(zego_error, unsigned long long msg_id, int);
extern void GoOnPlayerAudioData(unsigned char *, unsigned int, struct zego_audio_frame_param, char *);
extern void GoOnProcessRemoteAudioData(unsigned char *, unsigned int, struct zego_audio_frame_param *, char *, double);
extern void GoOnDebugError(int error_code, char* func_name, char* info);
extern void GoOnRoomStateUpdate(char *room_id, enum zego_room_state state, zego_error error_code, char *extend_data);
extern void GoOnRoomStreamUpdate(char *room_id, enum zego_update_type update_type, struct zego_stream *stream_info_list, unsigned int stream_info_count, char *extended_data);
extern void GoOnRoomStateChanged(char *room_id, enum zego_room_state_changed_reason reason, zego_error error_code, char *extended_data);
extern void GoOnRoomTokenWillExpire(char *room_id, int remain_time_in_second);
extern void GoOnPublisherStateUpdate(char *stream_id, enum zego_publisher_state state, zego_error error_code, char *extend_data);
extern void GoOnPublisherQualityUpdate(char *stream_id, struct zego_publish_stream_quality quality);
extern void GoOnPublisherStreamEvent(enum zego_stream_event event_id, char *stream_id, char *extra_info);
extern void GoOnPublisherSendAudioFirstFrame(enum zego_publish_channel channel);
extern void GoOnPlayerStateUpdate(char *stream_id, enum zego_player_state state, zego_error error_code, char *extend_data);
extern void GoOnPlayerQualityUpdate(char *stream_id, struct zego_play_stream_quality quality);
extern void GoOnPlayerRecvSei(struct zego_media_side_info info);
extern void GoOnPlayerStreamEvent(enum zego_stream_event event_id, char *stream_id, char *extra_info);
extern void GoOnPlayerRecvAudioFirstFrame(char *stream_id);

extern void GoOnMediaPlayerStateUpdate(enum zego_media_player_state state, zego_error error_code, enum zego_media_player_instance_index instance_index);
extern void GoOnMediaPlayerNetworkEvent(enum zego_media_player_network_event event, enum zego_media_player_instance_index instance_index);
extern void GoOnMediaPlayerPlayingProgress(unsigned long long millisecond, enum zego_media_player_instance_index instance_index);
extern void GoOnMediaPlayerRenderingProgress(unsigned long long millisecond, enum zego_media_player_instance_index instance_index);
extern void GoOnMediaPlayerRecvSEI(unsigned char *data, unsigned int data_length, enum zego_media_player_instance_index instance_index);
extern void GoOnMediaPlayerFirstFrameEvent(enum zego_media_player_first_frame_event event, enum zego_media_player_instance_index instance_index);
extern void GoOnMediaPlayerAudioFrame(unsigned char *data, unsigned int data_length, const struct zego_audio_frame_param param, enum zego_media_player_instance_index instance_index);
extern void GoOnMediaPlayerLoadFileResult(zego_error error_code, enum zego_media_player_instance_index instance_index);
extern void GoOnMediaPlayerSeekTo(zego_seq seq, zego_error error_code, enum zego_media_player_instance_index instance_index);
extern void GoOnEngineUninit();

static void bridge_go_on_api_called_result(int error_code, const char *func_name, const char *info, void *user_context) {
    GoOnApiCalledResult(error_code, (char *)func_name, (char *)info);
}

static void bridge_go_login_callback(zego_error code, const char *ext_data, const char *room_id, zego_seq seq, void *ctx) {
    GoLoginResultCallback(code, (char *)ext_data, seq);
}

static void bridge_go_logout_callback(zego_error code, const char *ext_data, const char *room_id, zego_seq seq, void *ctx) {
    GoLogoutResultCallback(code, (char *)ext_data, seq);
}

static void bridge_go_on_im_send_broadcast_message_result(const char *room_id, unsigned long long message_id, zego_error error_code, zego_seq seq, void *user_context) {
    GoOnIMSendBroadcastMessageResult(error_code, message_id, seq);
}

static void bridge_go_on_player_audio_data(const unsigned char *data, unsigned int data_length, struct zego_audio_frame_param param, const char *stream_id, void *user_context) {
    GoOnPlayerAudioData((unsigned char *)data, data_length, param, (char *)stream_id);
}

static void bridge_go_on_process_remote_audio_data(unsigned char *data, unsigned int data_length, struct zego_audio_frame_param *param, const char *stream_id, double timestamp, void *user_context) {
    GoOnProcessRemoteAudioData((unsigned char *)data, data_length, param, (char *)stream_id, timestamp);
}

static void bridge_go_on_debug_error(int error_code, const char *func_name, const char *info, void *user_context) {
    GoOnDebugError(error_code, (char *)func_name, (char *)info);
}

static void bridge_go_on_room_state_update(const char *room_id, enum zego_room_state state, zego_error error_code, const char *extend_data, void *user_context) {
    GoOnRoomStateUpdate((char *)room_id, state, error_code, (char *)extend_data);
}

static void bridge_go_on_room_stream_update(const char *room_id, enum zego_update_type update_type, const struct zego_stream *stream_info_list, unsigned int stream_info_count, const char *extended_data, void *user_context) {
    GoOnRoomStreamUpdate((char *)room_id, update_type, (struct zego_stream *)stream_info_list, stream_info_count, (char *)extended_data);
}

static void bridge_go_on_room_state_changed(const char *room_id, enum zego_room_state_changed_reason reason, zego_error error_code, const char *extended_data, void *user_context) {
    GoOnRoomStateChanged((char *)room_id, reason, error_code, (char *)extended_data);
}

static void bridge_go_on_room_token_will_expire(const char *room_id, int remain_time_in_second, void *user_context) {
    GoOnRoomTokenWillExpire((char *)room_id, remain_time_in_second);
}

static void bridge_go_on_publisher_state_update(const char *stream_id, enum zego_publisher_state state, zego_error error_code, const char *extend_data, void *user_context) {
    GoOnPublisherStateUpdate((char *)stream_id, state, error_code, (char *)extend_data);
}

static void bridge_go_on_publisher_quality_update(const char *stream_id, struct zego_publish_stream_quality quality, void *user_context) {
    GoOnPublisherQualityUpdate((char *)stream_id, quality);
}

static void bridge_go_on_publisher_stream_event(enum zego_stream_event event_id, const char *stream_id, const char *extra_info, void *user_context) {
    GoOnPublisherStreamEvent(event_id, (char *)stream_id, (char *)extra_info);
}

static void bridge_go_on_publisher_send_audio_first_frame(enum zego_publish_channel channel, void *user_context) {
    GoOnPublisherSendAudioFirstFrame(channel);
}

static void bridge_go_on_player_state_update(const char *stream_id, enum zego_player_state state, zego_error error_code, const char *extend_data, void *user_context) {
    GoOnPlayerStateUpdate((char *)stream_id, state, error_code, (char *)extend_data);
}

static void bridge_go_on_player_quality_update(const char *stream_id, struct zego_play_stream_quality quality, void *user_context) {
    GoOnPlayerQualityUpdate((char *)stream_id, quality);
}

static void bridge_go_on_player_recv_sei(struct zego_media_side_info info, void *user_context) {
    GoOnPlayerRecvSei(info);
}

static void bridge_go_on_player_stream_event(enum zego_stream_event event_id, const char *stream_id, const char *extra_info, void *user_context) {
    GoOnPlayerStreamEvent(event_id, (char *)stream_id, (char *)extra_info);
}

static void bridge_go_on_player_recv_audio_first_frame(const char *stream_id, void *user_context) {
    GoOnPlayerRecvAudioFirstFrame((char *)stream_id);
}

static void bridge_go_on_media_player_state_update(enum zego_media_player_state state, zego_error error_code, enum zego_media_player_instance_index instance_index, void *user_context) {
    GoOnMediaPlayerStateUpdate(state, error_code, instance_index);
}

static void bridge_go_on_media_player_network_event(enum zego_media_player_network_event event, enum zego_media_player_instance_index instance_index, void *user_context) {
    GoOnMediaPlayerNetworkEvent(event, instance_index);
}

static void bridge_go_on_media_player_playing_progress(unsigned long long millisecond, enum zego_media_player_instance_index instance_index, void *user_context) {
    GoOnMediaPlayerPlayingProgress(millisecond, instance_index);
}

static void bridge_go_on_media_player_rendering_progress(unsigned long long millisecond, enum zego_media_player_instance_index instance_index, void *user_context) {
    GoOnMediaPlayerRenderingProgress(millisecond, instance_index);
}

static void bridge_go_on_media_player_recv_sei(const unsigned char *data, unsigned int data_length, enum zego_media_player_instance_index instance_index, void *user_context) {
    GoOnMediaPlayerRecvSEI((char *)data, data_length, instance_index);
}

static void bridge_go_on_media_player_first_frame_event(enum zego_media_player_first_frame_event event, enum zego_media_player_instance_index instance_index, void *user_context) {
    GoOnMediaPlayerFirstFrameEvent(event, instance_index);
}

static void bridge_go_on_media_player_audio_frame(unsigned char *data, unsigned int data_length, const struct zego_audio_frame_param param, enum zego_media_player_instance_index instance_index, void *user_context) {
    GoOnMediaPlayerAudioFrame((char *)data, data_length, param, instance_index);
}

static void bridge_go_on_mediaplayer_load_file_result(zego_error error_code, enum zego_media_player_instance_index instance_index, void *user_context) {
    GoOnMediaPlayerLoadFileResult(error_code, instance_index);
}

static void bridge_go_on_media_player_seek_to(zego_seq seq, zego_error error_code, enum zego_media_player_instance_index instance_index, void *user_context) {
    GoOnMediaPlayerSeekTo(seq, error_code, instance_index);
}

static void bridge_go_on_engine_uninit(void *user_context) {
	GoOnEngineUninit();
}

static void zego_express_go_bridge_init() {
    zego_register_api_called_result_callback(bridge_go_on_api_called_result, NULL);
    zego_register_room_login_result_callback(bridge_go_login_callback, NULL);
    zego_register_room_logout_result_callback(bridge_go_logout_callback, NULL);
    zego_register_im_send_broadcast_message_result_callback(bridge_go_on_im_send_broadcast_message_result, NULL);
    zego_register_player_audio_data_callback(bridge_go_on_player_audio_data, NULL);
		zego_register_process_remote_audio_data_callback(bridge_go_on_process_remote_audio_data, NULL);
    zego_register_debug_error_callback(bridge_go_on_debug_error, NULL);
    zego_register_room_state_update_callback(bridge_go_on_room_state_update, NULL);
    zego_register_room_stream_update_callback(bridge_go_on_room_stream_update, NULL);
    zego_register_room_state_changed_callback(bridge_go_on_room_state_changed, NULL);
    zego_register_room_token_will_expire_callback(bridge_go_on_room_token_will_expire, NULL);
    zego_register_publisher_state_update_callback(bridge_go_on_publisher_state_update, NULL);
    zego_register_publisher_quality_update_callback(bridge_go_on_publisher_quality_update, NULL);
    zego_register_publisher_stream_event_callback(bridge_go_on_publisher_stream_event, NULL);
    zego_register_publisher_send_audio_first_frame_callback(bridge_go_on_publisher_send_audio_first_frame, NULL);
    zego_register_player_state_update_callback(bridge_go_on_player_state_update, NULL);
    zego_register_player_quality_update_callback(bridge_go_on_player_quality_update, NULL);
    zego_register_player_recv_media_side_info_callback(bridge_go_on_player_recv_sei, NULL);
    zego_register_player_stream_event_callback(bridge_go_on_player_stream_event, NULL);
    zego_register_player_recv_audio_first_frame_callback(bridge_go_on_player_recv_audio_first_frame, NULL);
    zego_register_media_player_state_update_callback(bridge_go_on_media_player_state_update, NULL);
    zego_register_media_player_network_event_callback(bridge_go_on_media_player_network_event, NULL);
    zego_register_media_player_playing_progress_callback(bridge_go_on_media_player_playing_progress, NULL);
    zego_register_media_player_rendering_progress_callback(bridge_go_on_media_player_rendering_progress, NULL);
    zego_register_media_player_recv_sei_callback(bridge_go_on_media_player_recv_sei, NULL);
    zego_register_media_player_first_frame_event_callback(bridge_go_on_media_player_first_frame_event, NULL);
    zego_register_media_player_audio_frame_callback(bridge_go_on_media_player_audio_frame, NULL);
    zego_register_media_player_load_resource_callback(bridge_go_on_mediaplayer_load_file_result, NULL);
    zego_register_media_player_seek_to_callback(bridge_go_on_media_player_seek_to, NULL);
    zego_register_engine_uninit_callback(bridge_go_on_engine_uninit, NULL);
}

*/
import "C"
import (
	"container/list"
	"fmt"
	"strconv"
	"sync"
	"unsafe"
)

func init() {
	fmt.Println("zegoexpress init")
	C.zego_express_go_bridge_init()
}

var (
	engineLock                sync.RWMutex
	globalEngine              *engineImpl
	engineDestroyCallbackLock sync.Mutex
	engineDestroyCallback     ZegoDestroyCompletionCallback
	maxPublishChannelCount    int = 4

	callbackLock                   sync.Mutex
	apiCalledCallback              IZegoApiCalledEventHandler
	roomLoginCallback              = make(map[int]ZegoRoomLoginCallback)
	roomLogoutCallback             = make(map[int]ZegoRoomLogoutCallback)
	imSendBroadcastMessageCallback = make(map[int]ZegoIMSendBroadcastMessageCallback)

	mediaPlayerLock    sync.Mutex
	mediaPlayerImplMap = make(map[int]*mediaPlayerImpl)
)

//export GoOnApiCalledResult
func GoOnApiCalledResult(errorCode C.int, funcName *C.char, data *C.char) {
	callbackLock.Lock()
	defer callbackLock.Unlock()
	if apiCalledCallback != nil {
		goData := ""
		if data != nil {
			goData = C.GoString(data)
		}
		apiCalledCallback.OnApiCalledResult(int(errorCode), C.GoString(funcName), goData)
	}
}

//export GoLoginResultCallback
func GoLoginResultCallback(errorCode C.zego_error, extendedData *C.char, seq C.zego_seq) {
	callbackLock.Lock()
	defer callbackLock.Unlock()

	callback, ok := roomLoginCallback[int(seq)]
	if !ok {
		return
	}

	if callback != nil {
		goExtendedData := ""
		if extendedData != nil {
			goExtendedData = C.GoString(extendedData)
		}
		callback(int(errorCode), goExtendedData)
	}

	delete(roomLoginCallback, int(seq))
}

//export GoLogoutResultCallback
func GoLogoutResultCallback(errorCode C.zego_error, extendedData *C.char, seq C.zego_seq) {
	callbackLock.Lock()
	defer callbackLock.Unlock()

	callback, ok := roomLogoutCallback[int(seq)]
	if !ok {
		return
	}

	if callback != nil {
		goExtendedData := ""
		if extendedData != nil {
			goExtendedData = C.GoString(extendedData)
		}
		callback(int(errorCode), goExtendedData)
	}

	delete(roomLogoutCallback, int(seq))
}

//export GoOnIMSendBroadcastMessageResult
func GoOnIMSendBroadcastMessageResult(errorCode C.zego_error, messageID C.ulonglong, seq C.zego_seq) {
	callbackLock.Lock()
	defer callbackLock.Unlock()

	callback, ok := imSendBroadcastMessageCallback[int(seq)]
	if !ok {
		return
	}

	if callback != nil {
		callback(int(errorCode), uint64(messageID))
	}

	delete(imSendBroadcastMessageCallback, int(seq))
}

//export GoOnPlayerAudioData
func GoOnPlayerAudioData(data *C.uchar, dataLen C.uint, param C.struct_zego_audio_frame_param, streamID *C.char) {
	engineLock.RLock()
	defer engineLock.RUnlock()
	if globalEngine == nil {
		return
	}
	handler := globalEngine.audioDataHandler
	if handler == nil {
		return
	}
	goData := cUcharPtrToGoSlice(data, dataLen)
	goParam := ZegoAudioFrameParam{
		SampleRate: ZegoAudioSampleRate(param.sample_rate),
		Channel:    ZegoAudioChannel(param.channel),
	}
	handler.OnPlayerAudioData(goData, goParam, C.GoString(streamID))
}

//export GoOnProcessRemoteAudioData
func GoOnProcessRemoteAudioData(data *C.uchar, dataLen C.uint, param *C.struct_zego_audio_frame_param, streamID *C.char, timestamp float64) {
	engineLock.RLock()
	defer engineLock.RUnlock()
	if globalEngine == nil {
		return
	}
	handler := globalEngine.customAudioProcessHandler
	if handler == nil {
		return
	}
	goData := cUcharPtrToGoSlice(data, dataLen)
	goParam := ZegoAudioFrameParam{
		SampleRate: ZegoAudioSampleRate(param.sample_rate),
		Channel:    ZegoAudioChannel(param.channel),
	}
	handler.OnProcessRemoteAudioData(goData, &goParam, C.GoString(streamID), timestamp)
	param.sample_rate = C.enum_zego_audio_sample_rate(goParam.SampleRate)
	param.channel = C.enum_zego_audio_channel(goParam.Channel)
}

//export GoOnDebugError
func GoOnDebugError(errorCode C.int, funcName *C.char, info *C.char) {
	engineLock.RLock()
	defer engineLock.RUnlock()
	if globalEngine == nil {
		return
	}
	handler := globalEngine.eventHandler
	if handler == nil {
		return
	}
	goInfo := ""
	if info != nil {
		goInfo = C.GoString(info)
	}
	handler.OnDebugError(int(errorCode), C.GoString(funcName), goInfo)
}

//export GoOnRoomStateUpdate
func GoOnRoomStateUpdate(roomID *C.char, state C.enum_zego_room_state, errorCode C.zego_error, data *C.char) {
	engineLock.RLock()
	defer engineLock.RUnlock()
	if globalEngine == nil {
		return
	}
	handler := globalEngine.eventHandler
	if handler == nil {
		return
	}
	goData := ""
	if data != nil {
		goData = C.GoString(data)
	}
	handler.OnRoomStateUpdate(C.GoString(roomID), ZegoRoomState(state), int(errorCode), goData)
}

//export GoOnRoomStreamUpdate
func GoOnRoomStreamUpdate(roomID *C.char, updateType C.enum_zego_update_type, streamInfoList *C.struct_zego_stream, streamInfoCount C.uint, data *C.char) {
	engineLock.RLock()
	defer engineLock.RUnlock()
	if globalEngine == nil {
		return
	}
	handler := globalEngine.eventHandler
	if handler == nil {
		return
	}
	streamList := make([]ZegoStream, 0)
	if streamInfoList != nil && streamInfoCount > 0 {
		cStreams := unsafe.Slice(streamInfoList, streamInfoCount)

		for i := 0; i < int(streamInfoCount); i++ {
			stream := cStreams[i]
			streamList = append(streamList, convertStream(stream))
		}
	}
	goData := ""
	if data != nil {
		goData = C.GoString(data)
	}
	handler.OnRoomStreamUpdate(C.GoString(roomID), ZegoUpdateType(updateType), streamList, goData)
}

//export GoOnRoomStateChanged
func GoOnRoomStateChanged(roomID *C.char, reason C.enum_zego_room_state_changed_reason, errorCode C.zego_error, data *C.char) {
	engineLock.RLock()
	defer engineLock.RUnlock()
	if globalEngine == nil {
		return
	}
	handler := globalEngine.eventHandler
	if handler == nil {
		return
	}
	goData := ""
	if data != nil {
		goData = C.GoString(data)
	}
	handler.OnRoomStateChanged(C.GoString(roomID), ZegoRoomStateChangedReason(reason), int(errorCode), goData)
}

//export GoOnRoomTokenWillExpire
func GoOnRoomTokenWillExpire(roomID *C.char, remainTimeInSecond C.int) {
	engineLock.RLock()
	defer engineLock.RUnlock()
	if globalEngine == nil {
		return
	}
	handler := globalEngine.eventHandler
	if handler == nil {
		return
	}
	handler.OnRoomTokenWillExpire(C.GoString(roomID), int(remainTimeInSecond))
}

//export GoOnPublisherStateUpdate
func GoOnPublisherStateUpdate(streamID *C.char, state C.enum_zego_publisher_state, errorCode C.zego_error, data *C.char) {
	engineLock.RLock()
	defer engineLock.RUnlock()
	if globalEngine == nil {
		return
	}
	handler := globalEngine.eventHandler
	if handler == nil {
		return
	}
	goData := ""
	if data != nil {
		goData = C.GoString(data)
	}
	handler.OnPublisherStateUpdate(C.GoString(streamID), ZegoPublisherState(state), int(errorCode), goData)
}

//export GoOnPublisherQualityUpdate
func GoOnPublisherQualityUpdate(streamID *C.char, quality C.struct_zego_publish_stream_quality) {
	engineLock.RLock()
	defer engineLock.RUnlock()
	if globalEngine == nil {
		return
	}
	handler := globalEngine.eventHandler
	if handler == nil {
		return
	}
	goQuality := ZegoPublishStreamQuality{
		VideoCaptureFPS:  float64(quality.video_capture_fps),
		VideoEncodeFPS:   float64(quality.video_encode_fps),
		VideoSendFPS:     float64(quality.video_send_fps),
		VideoKBPS:        float64(quality.video_kbps),
		AudioCaptureFPS:  float64(quality.audio_capture_fps),
		AudioSendFPS:     float64(quality.audio_send_fps),
		AudioKBPS:        float64(quality.audio_kbps),
		Rtt:              int(quality.rtt),
		PacketLostRate:   float64(quality.packet_lost_rate),
		Level:            ZegoStreamQualityLevel(quality.level),
		IsHardwareEncode: bool(quality.is_hardware_encode),
		VideoCodecID:     ZegoVideoCodecID(quality.video_codec_id),
		TotalSendBytes:   float64(quality.total_send_bytes),
		AudioSendBytes:   float64(quality.audio_send_bytes),
		VideoSendBytes:   float64(quality.video_send_bytes),
	}
	handler.OnPublisherQualityUpdate(C.GoString(streamID), goQuality)
}

//export GoOnPublisherStreamEvent
func GoOnPublisherStreamEvent(eventID C.enum_zego_stream_event, streamID *C.char, data *C.char) {
	engineLock.RLock()
	defer engineLock.RUnlock()
	if globalEngine == nil {
		return
	}
	handler := globalEngine.eventHandler
	if handler == nil {
		return
	}
	goData := ""
	if data != nil {
		goData = C.GoString(data)
	}
	handler.OnPublisherStreamEvent(ZegoStreamEvent(eventID), C.GoString(streamID), goData)
}

//export GoOnPublisherSendAudioFirstFrame
func GoOnPublisherSendAudioFirstFrame(channel C.enum_zego_publish_channel) {
	engineLock.RLock()
	defer engineLock.RUnlock()
	if globalEngine == nil {
		return
	}
	handler := globalEngine.eventHandler
	if handler == nil {
		return
	}
	handler.OnPublisherSendAudioFirstFrame(ZegoPublishChannel(channel))
}

//export GoOnPlayerStateUpdate
func GoOnPlayerStateUpdate(streamID *C.char, state C.enum_zego_player_state, errorCode C.zego_error, data *C.char) {
	engineLock.RLock()
	defer engineLock.RUnlock()
	if globalEngine == nil {
		return
	}
	handler := globalEngine.eventHandler
	if handler == nil {
		return
	}
	goData := ""
	if data != nil {
		goData = C.GoString(data)
	}
	handler.OnPlayerStateUpdate(C.GoString(streamID), ZegoPlayerState(state), int(errorCode), goData)
}

//export GoOnPlayerQualityUpdate
func GoOnPlayerQualityUpdate(streamID *C.char, quality C.struct_zego_play_stream_quality) {
	engineLock.RLock()
	defer engineLock.RUnlock()
	if globalEngine == nil {
		return
	}
	handler := globalEngine.eventHandler
	if handler == nil {
		return
	}
	goQuality := ZegoPlayStreamQuality{
		VideoRecvFPS:              float64(quality.video_recv_fps),
		VideoDejitterFPS:          float64(quality.video_dejitter_fps),
		VideoDecodeFPS:            float64(quality.video_decode_fps),
		VideoRenderFPS:            float64(quality.video_render_fps),
		VideoKBPS:                 float64(quality.video_kbps),
		VideoBreakRate:            float64(quality.video_break_rate),
		AudioRecvFPS:              float64(quality.audio_recv_fps),
		AudioDejitterFPS:          float64(quality.audio_dejitter_fps),
		AudioDecodeFPS:            float64(quality.audio_decode_fps),
		AudioRenderFPS:            float64(quality.audio_render_fps),
		AudioKBPS:                 float64(quality.audio_kbps),
		AudioBreakRate:            float64(quality.audio_break_rate),
		Mos:                       float64(quality.mos),
		Rtt:                       int(quality.rtt),
		PacketLostRate:            float64(quality.packet_lost_rate),
		PeerToPeerDelay:           int(quality.peer_to_peer_delay),
		PeerToPeerPacketLostRate:  float64(quality.peer_to_peer_packet_lost_rate),
		Level:                     ZegoStreamQualityLevel(quality.level),
		Delay:                     int(quality.delay),
		AvTimestampDiff:           int(quality.av_timestamp_diff),
		IsHardwareDecode:          bool(quality.is_hardware_decode),
		VideoCodecID:              ZegoVideoCodecID(quality.video_codec_id),
		TotalRecvBytes:            float64(quality.total_recv_bytes),
		AudioRecvBytes:            float64(quality.audio_recv_bytes),
		VideoRecvBytes:            float64(quality.video_recv_bytes),
		AudioCumulativeBreakCount: uint32(quality.audio_cumulative_break_count),
		AudioCumulativeBreakTime:  uint32(quality.audio_cumulative_break_time),
		AudioCumulativeBreakRate:  float64(quality.audio_cumulative_break_rate),
		AudioCumulativeDecodeTime: uint32(quality.audio_cumulative_decode_time),
		VideoCumulativeBreakCount: uint32(quality.video_cumulative_break_count),
		VideoCumulativeBreakTime:  uint32(quality.video_cumulative_break_time),
		VideoCumulativeBreakRate:  float64(quality.video_cumulative_break_rate),
		VideoCumulativeDecodeTime: uint32(quality.video_cumulative_decode_time),
		MuteVideo:                 int(quality.mute_video),
		MuteAudio:                 int(quality.mute_audio),
	}
	handler.OnPlayerQualityUpdate(C.GoString(streamID), goQuality)
}

//export GoOnPlayerRecvSei
func GoOnPlayerRecvSei(info C.struct_zego_media_side_info) {
	engineLock.RLock()
	defer engineLock.RUnlock()
	if globalEngine == nil {
		return
	}
	handler := globalEngine.eventHandler
	if handler == nil {
		return
	}
	goInfo := ZegoMediaSideInfo{
		StreamID:    C.GoString(&info.stream_id[0]),
		SEIData:     cUcharPtrToGoSlice(info.sei_data, info.sei_data_length),
		TimestampNs: int64(info.timestamp_ns),
		ModuleType:  int(info.module_type),
	}
	handler.OnPlayerRecvMediaSideInfo(goInfo)
}

//export GoOnPlayerStreamEvent
func GoOnPlayerStreamEvent(eventID C.enum_zego_stream_event, streamID *C.char, data *C.char) {
	engineLock.RLock()
	defer engineLock.RUnlock()
	if globalEngine == nil {
		return
	}
	handler := globalEngine.eventHandler
	if handler == nil {
		return
	}
	goData := ""
	if data != nil {
		goData = C.GoString(data)
	}
	handler.OnPlayerStreamEvent(ZegoStreamEvent(eventID), C.GoString(streamID), goData)
}

//export GoOnPlayerRecvAudioFirstFrame
func GoOnPlayerRecvAudioFirstFrame(streamID *C.char) {
	engineLock.RLock()
	defer engineLock.RUnlock()
	if globalEngine == nil {
		return
	}
	handler := globalEngine.eventHandler
	if handler == nil {
		return
	}
	handler.OnPlayerRecvAudioFirstFrame(C.GoString(streamID))
}

//export GoOnMediaPlayerStateUpdate
func GoOnMediaPlayerStateUpdate(state C.enum_zego_media_player_state, errorCode C.zego_error, index C.enum_zego_media_player_instance_index) {
	mediaPlayerLock.Lock()
	defer mediaPlayerLock.Unlock()
	if mediaPlayer, ok := mediaPlayerImplMap[int(index)]; ok {
		handler := mediaPlayer.eventHandler
		if handler != nil {
			handler.OnMediaPlayerStateUpdate(mediaPlayer, ZegoMediaPlayerState(state), int(errorCode))
		}
	}
}

//export GoOnMediaPlayerNetworkEvent
func GoOnMediaPlayerNetworkEvent(event C.enum_zego_media_player_network_event, index C.enum_zego_media_player_instance_index) {
	mediaPlayerLock.Lock()
	defer mediaPlayerLock.Unlock()
	if mediaPlayer, ok := mediaPlayerImplMap[int(index)]; ok {
		handler := mediaPlayer.eventHandler
		if handler != nil {
			handler.OnMediaPlayerNetworkEvent(mediaPlayer, ZegoMediaPlayerNetworkEvent(event))
		}
	}
}

//export GoOnMediaPlayerPlayingProgress
func GoOnMediaPlayerPlayingProgress(millisecond C.ulonglong, index C.enum_zego_media_player_instance_index) {
	mediaPlayerLock.Lock()
	defer mediaPlayerLock.Unlock()
	if mediaPlayer, ok := mediaPlayerImplMap[int(index)]; ok {
		handler := mediaPlayer.eventHandler
		if handler != nil {
			handler.OnMediaPlayerPlayingProgress(mediaPlayer, uint64(millisecond))
		}
	}
}

//export GoOnMediaPlayerRenderingProgress
func GoOnMediaPlayerRenderingProgress(millisecond C.ulonglong, index C.enum_zego_media_player_instance_index) {
	mediaPlayerLock.Lock()
	defer mediaPlayerLock.Unlock()
	if mediaPlayer, ok := mediaPlayerImplMap[int(index)]; ok {
		handler := mediaPlayer.eventHandler
		if handler != nil {
			handler.OnMediaPlayerRenderingProgress(mediaPlayer, uint64(millisecond))
		}
	}
}

//export GoOnMediaPlayerRecvSEI
func GoOnMediaPlayerRecvSEI(data *C.uchar, dataLen C.uint, index C.enum_zego_media_player_instance_index) {
	mediaPlayerLock.Lock()
	defer mediaPlayerLock.Unlock()
	if mediaPlayer, ok := mediaPlayerImplMap[int(index)]; ok {
		handler := mediaPlayer.eventHandler
		if handler != nil {
			goData := cUcharPtrToGoSlice(data, dataLen)
			handler.OnMediaPlayerRecvSEI(mediaPlayer, goData)
		}
	}
}

//export GoOnMediaPlayerFirstFrameEvent
func GoOnMediaPlayerFirstFrameEvent(event C.enum_zego_media_player_first_frame_event, index C.enum_zego_media_player_instance_index) {
	mediaPlayerLock.Lock()
	defer mediaPlayerLock.Unlock()
	if mediaPlayer, ok := mediaPlayerImplMap[int(index)]; ok {
		handler := mediaPlayer.eventHandler
		if handler != nil {
			handler.OnMediaPlayerFirstFrameEvent(mediaPlayer, ZegoMediaPlayerFirstFrameEvent(event))
		}
	}
}

//export GoOnMediaPlayerAudioFrame
func GoOnMediaPlayerAudioFrame(data *C.uchar, dataLen C.uint, param C.struct_zego_audio_frame_param, index C.enum_zego_media_player_instance_index) {
	mediaPlayerLock.Lock()
	defer mediaPlayerLock.Unlock()
	if mediaPlayer, ok := mediaPlayerImplMap[int(index)]; ok {
		handler := mediaPlayer.audioHandler
		if handler != nil {
			goData := cUcharPtrToGoSlice(data, dataLen)
			goParam := ZegoAudioFrameParam{
				SampleRate: ZegoAudioSampleRate(param.sample_rate),
				Channel:    ZegoAudioChannel(param.channel),
			}
			handler.OnAudioFrame(mediaPlayer, goData, goParam)
		}
	}
}

//export GoOnMediaPlayerLoadFileResult
func GoOnMediaPlayerLoadFileResult(errorCode C.zego_error, index C.enum_zego_media_player_instance_index) {
	mediaPlayerLock.Lock()
	defer mediaPlayerLock.Unlock()
	if mediaPlayer, ok := mediaPlayerImplMap[int(index)]; ok {
		callbacks := mediaPlayer.loadResourceCallbacks
		if callbacks.Len() > 0 {
			callback := callbacks.Front().Value
			callbacks.Remove(callbacks.Front())
			if callback != nil {
				if f, ok := callback.(ZegoMediaPlayerLoadResourceCallback); ok {
					f(int(errorCode))
				}
			}
		}
	}
}

//export GoOnMediaPlayerSeekTo
func GoOnMediaPlayerSeekTo(seq C.zego_seq, errorCode C.zego_error, index C.enum_zego_media_player_instance_index) {
	mediaPlayerLock.Lock()
	defer mediaPlayerLock.Unlock()
	if mediaPlayer, ok := mediaPlayerImplMap[int(index)]; ok {
		callbacks := mediaPlayer.seekToCallbacks
		if callback, ok := callbacks[int(seq)]; ok {
			if callback != nil {
				callback(int(errorCode))
			}
			delete(callbacks, int(seq))
		}
	}
}

//export GoOnEngineUninit
func GoOnEngineUninit() {
	engineDestroyCallbackLock.Lock()
	defer engineDestroyCallbackLock.Unlock()
	if engineDestroyCallback != nil {
		engineDestroyCallback()
		engineDestroyCallback = nil
	}
}

type engineImpl struct {
	eventHandler              IZegoEventHandler
	audioDataHandler          IZegoAudioDataHandler
	customAudioProcessHandler IZegoCustomAudioProcessHandler
}

func (e *engineImpl) init(profile ZegoEngineProfile, handler IZegoEventHandler) int {
	e.eventHandler = handler
	engineConfig := ZegoEngineConfig{
		LogConfig: nil,
		AdvancedConfig: map[string]string{
			"thirdparty_framework_info": "golang",
		},
	}
	setEngineConfig(engineConfig)
	cProfile := C.struct_zego_engine_profile{
		app_id:   C.uint(profile.AppID),
		scenario: C.enum_zego_scenario(profile.Scenario),
	}
	setCharArray(&cProfile.app_sign[0], profile.AppSign, C.ZEGO_EXPRESS_MAX_APPSIGN_LEN)
	return int(C.zego_express_engine_init_with_profile(cProfile))
}

func (e *engineImpl) EnableDebugAssistant(enable bool) {
	C.zego_express_enable_debug_assistant(C.bool(enable))
}

func (e *engineImpl) LoginRoom(roomID string, user ZegoUser, config *ZegoRoomConfig, callback ZegoRoomLoginCallback) {
	var seq C.int
	cRoomID := StringToCString(roomID)
	defer FreeCString(cRoomID)
	var cZegoUser C.struct_zego_user
	setCharArray(&cZegoUser.user_id[0], user.UserID, C.ZEGO_EXPRESS_MAX_USERID_LEN)
	setCharArray(&cZegoUser.user_name[0], user.UserName, C.ZEGO_EXPRESS_MAX_USERNAME_LEN)
	var cRoomConfig C.struct_zego_room_config
	var cRoomConfigPtr *C.struct_zego_room_config = nil
	if config != nil {
		cRoomConfig.max_member_count = C.uint(config.MaxMemberCount)
		cRoomConfig.is_user_status_notify = C.bool(config.IsUserStatusNotify)
		setCharArray(&cRoomConfig.token[0], config.Token, C.ZEGO_EXPRESS_MAX_ROOM_TOKEN_VALUE_LEN)
		cRoomConfig.capability_negotiation_types = C.uint(config.CapabilityNegotiationTypes)
		cRoomConfig.room_type = C.uint(config.RoomType)
		cRoomConfigPtr = &cRoomConfig
	}
	C.zego_express_login_room_with_callback(cRoomID, cZegoUser, cRoomConfigPtr, &seq)
	if callback != nil {
		callbackLock.Lock()
		defer callbackLock.Unlock()
		roomLoginCallback[int(seq)] = callback
	}
}

func (e *engineImpl) LogoutRoom(roomID string, callback ZegoRoomLogoutCallback) {
	var seq C.int
	cRoomID := StringToCString(roomID)
	defer FreeCString(cRoomID)
	C.zego_express_logout_room_with_callback(cRoomID, &seq)
	if callback != nil {
		callbackLock.Lock()
		defer callbackLock.Unlock()
		roomLogoutCallback[int(seq)] = callback
	}
}

func (e *engineImpl) RenewToken(roomID string, token string) {
	cRoomID := StringToCString(roomID)
	defer FreeCString(cRoomID)
	cToken := StringToCString(token)
	defer FreeCString(cToken)
	C.zego_express_renew_token(cRoomID, cToken)
}

func (e *engineImpl) SendBroadcastMessage(roomID string, message string, callback ZegoIMSendBroadcastMessageCallback) {
	var seq C.int

	cRoomID := StringToCString(roomID)
	defer FreeCString(cRoomID)
	cMessage := StringToCString(message)
	defer FreeCString(cMessage)

	C.zego_express_send_broadcast_message(cRoomID, cMessage, &seq)
	if callback != nil {
		callbackLock.Lock()
		defer callbackLock.Unlock()
		imSendBroadcastMessageCallback[int(seq)] = callback
	}
}

func (e *engineImpl) GetRoomStreamList(roomID string, streamListType ZegoRoomStreamListType) ZegoRoomStreamList {
	result := ZegoRoomStreamList{
		PublishStreamList: make([]ZegoStream, 0),
		PlayStreamList:    make([]ZegoStream, 0),
	}
	cRoomID := StringToCString(roomID)
	defer FreeCString(cRoomID)
	var cResult *C.struct_zego_room_stream_list
	C.zego_express_get_room_stream_list(cRoomID, C.enum_zego_room_stream_list_type(streamListType), &cResult)
	if cResult.publish_stream_list != nil && cResult.publish_stream_list_count > 0 {
		streams := unsafe.Slice(cResult.publish_stream_list, cResult.publish_stream_list_count)

		for i := 0; i < int(cResult.publish_stream_list_count); i++ {
			stream := streams[i]
			result.PublishStreamList = append(result.PublishStreamList, convertStream(stream))
		}
	}
	if cResult.play_stream_list != nil && cResult.play_stream_list_count > 0 {
		streams := unsafe.Slice(cResult.play_stream_list, cResult.play_stream_list_count)

		for i := 0; i < int(cResult.play_stream_list_count); i++ {
			stream := streams[i]
			result.PlayStreamList = append(result.PlayStreamList, convertStream(stream))
		}
	}
	C.zego_express_free_room_stream_list(cResult)
	return result
}

func (e *engineImpl) StartPublishingStream(streamID string, config ZegoPublisherConfig, channel ZegoPublishChannel) {
	cConfig := C.struct_zego_publisher_config{
		force_synchronous_network_time: C.int(config.ForceSynchronousNetworkTime),
		stream_censorship_mode:         C.enum_zego_stream_censorship_mode(config.StreamCensorshipMode),
		stream_censor_flag:             C.int(config.StreamCensorFlag),
		codec_negotiation_type:         C.enum_zego_capability_negotiation_type(config.CodecNegotiationType),
	}
	setCharArray(&cConfig.room_id[0], config.RoomID, C.ZEGO_EXPRESS_MAX_ROOMID_LEN)
	setCharArray(&cConfig.stream_title[0], config.StreamTitle, C.ZEGO_EXPRESS_MAX_STREAM_TITLE_LEN)

	cStreamID := StringToCString(streamID)
	defer FreeCString(cStreamID)

	C.zego_express_start_publishing_stream_with_config(cStreamID, cConfig, C.enum_zego_publish_channel(channel))
}

func (e *engineImpl) StopPublishingStream(channel ZegoPublishChannel) {
	C.zego_express_stop_publishing_stream(C.enum_zego_publish_channel(channel))
}

func (e *engineImpl) SetAudioConfig(config ZegoAudioConfig, channel ZegoPublishChannel) {
	cConfig := C.struct_zego_audio_config{
		bitrate:  C.int(config.Bitrate),
		channel:  C.enum_zego_audio_channel(config.Channel),
		codec_id: C.enum_zego_audio_codec_id(config.CodecID),
	}
	C.zego_express_set_audio_config_by_channel(cConfig, C.enum_zego_publish_channel(channel))
}

func (e *engineImpl) EnableAEC(enable bool) {
	C.zego_express_enable_aec(C.bool(enable))
}

func (e *engineImpl) EnableCustomAudioIO(enable bool, config *ZegoCustomAudioConfig, channel ZegoPublishChannel) {
	cConfig := C.struct_zego_custom_audio_config{
		source_type: C.zego_audio_source_type_default,
	}
	if config != nil {
		cConfig.source_type = C.enum_zego_audio_source_type(config.SourceType)
	}
	C.zego_express_enable_custom_audio_io(C.bool(enable), &cConfig, C.enum_zego_publish_channel(channel))
}

func (e *engineImpl) SendSEI(data []uint8, channel ZegoPublishChannel) {
	cData, cLen := goSliceToCUchar(data)
	C.zego_express_send_sei(cData, cLen, C.enum_zego_publish_channel(channel))
}

func (e *engineImpl) SetAudioDataHandler(handler IZegoAudioDataHandler) {
	engineLock.Lock()
	defer engineLock.Unlock()
	e.audioDataHandler = handler
}

func (e *engineImpl) StartAudioDataObserver(observerBitMask uint32, param ZegoAudioFrameParam) {
	C.zego_express_start_audio_data_observer(C.uint(observerBitMask), convertAudioFrameParam(param))
}

func (e *engineImpl) StopAudioDataObserver() {
	C.zego_express_stop_audio_data_observer()
}

func (e *engineImpl) StartPlayingStream(streamID string, config *ZegoPlayerConfig) {
	cStreamID := StringToCString(streamID)
	defer FreeCString(cStreamID)
	if config == nil {
		C.zego_express_start_playing_stream(cStreamID, nil)
		return
	}

	var cCdnConfig C.struct_zego_cdn_config
	var cCdnConfigPtr *C.struct_zego_cdn_config = nil
	if config.CdnConfig != nil {
		setCharArray(&cCdnConfig.url[0], config.CdnConfig.Url, C.ZEGO_EXPRESS_MAX_URL_LEN)
		setCharArray(&cCdnConfig.auth_param[0], config.CdnConfig.AuthParam, C.ZEGO_EXPRESS_MAX_COMMON_LEN)
		setCharArray(&cCdnConfig.protocol[0], config.CdnConfig.Protocol, C.ZEGO_EXPRESS_MAX_COMMON_LEN)
		setCharArray(&cCdnConfig.quic_version[0], config.CdnConfig.QuicVersion, C.ZEGO_EXPRESS_MAX_COMMON_LEN)
		cCdnConfig.http_dns = C.enum_zego_http_dns_type(config.CdnConfig.Httpdns)
		cCdnConfig.quic_connect_mode = C.int(config.CdnConfig.QuicConnectMode)
		setCharArray(&cCdnConfig.custom_params[0], config.CdnConfig.CustomParam, C.ZEGO_EXPRESS_MAX_COMMON_LEN)
		cCdnConfigPtr = &cCdnConfig
	}
	cConfig := C.struct_zego_player_config{
		resource_mode:        C.enum_zego_stream_resource_mode(config.ResourceMode),
		cdn_config:           cCdnConfigPtr,
		video_codec_id:       C.zego_video_codec_id_unknown,
		source_resource_type: C.zego_resource_type_rtc,
	}
	setCharArray(&cConfig.room_id[0], config.RoomID, C.ZEGO_EXPRESS_MAX_ROOMID_LEN)
	C.zego_express_start_playing_stream_with_config(cStreamID, nil, cConfig)
}

func (e *engineImpl) StopPlayingStream(streamID string) {
	cStreamID := StringToCString(streamID)
	defer FreeCString(cStreamID)
	C.zego_express_stop_playing_stream(cStreamID)
}

func (e *engineImpl) SendCustomAudioCapturePCMData(data []uint8, param ZegoAudioFrameParam, channel ZegoPublishChannel) {
	cData, cLen := goSliceToCUchar(data)
	cParam := convertAudioFrameParam(param)
	C.zego_express_send_custom_audio_capture_pcm_data(cData, cLen, cParam, C.enum_zego_publish_channel(channel))
}

func (e *engineImpl) FetchCustomAudioRenderPCMData(data []uint8, param ZegoAudioFrameParam) {
	cData, cLen := goSliceToCUchar(data)
	cParam := convertAudioFrameParam(param)
	C.zego_express_fetch_custom_audio_render_pcm_data(cData, cLen, cParam)
}

func (e *engineImpl) SetCustomAudioProcessHandler(handle IZegoCustomAudioProcessHandler) {
	engineLock.Lock()
	defer engineLock.Unlock()
	e.customAudioProcessHandler = handle
}

func (e *engineImpl) EnableCustomAudioRemoteProcessing(enable bool, config *ZegoCustomAudioProcessConfig) {
	var cConfig C.struct_zego_custom_audio_process_config
	var cConfigPtr *C.struct_zego_custom_audio_process_config = nil
	if config != nil {
		cConfig.sample_rate = C.enum_zego_audio_sample_rate(config.SampleRate)
		cConfig.channel = C.enum_zego_audio_channel(config.Channel)
		cConfig.samples = C.int(config.Samples)
		cConfigPtr = &cConfig
	}
	C.zego_express_enable_custom_audio_remote_processing(C.bool(enable), cConfigPtr)
}

func (e *engineImpl) CreateMediaPlayer() IZegoMediaPlayer {
	var index C.enum_zego_media_player_instance_index = C.zego_media_player_instance_index_null
	C.zego_express_create_media_player(&index)
	if index == C.zego_media_player_instance_index_null {
		return nil
	}
	mediaPlayerLock.Lock()
	defer mediaPlayerLock.Unlock()
	mediaPlayer := new(mediaPlayerImpl)
	mediaPlayer.instanceIndex = int(index)
	mediaPlayer.loadResourceCallbacks = list.New()
	mediaPlayer.seekToCallbacks = make(map[int]ZegoMediaPlayerSeekToCallback)
	mediaPlayerImplMap[int(index)] = mediaPlayer
	return mediaPlayer
}

func (e *engineImpl) DestroyMediaPlayer(mediaPlayer IZegoMediaPlayer) {
	if mediaPlayer == nil {
		return
	}
	index := mediaPlayer.GetIndex()

	mediaPlayerLock.Lock()
	defer mediaPlayerLock.Unlock()
	if _, ok := mediaPlayerImplMap[index]; ok {
		C.zego_express_destroy_media_player(C.enum_zego_media_player_instance_index(index))
		delete(mediaPlayerImplMap, index)
	}
}

type mediaPlayerImpl struct {
	eventHandler          IZegoMediaPlayerEventHandler
	audioHandler          IZegoMediaPlayerAudioHandler
	instanceIndex         int
	loadResourceCallbacks *list.List
	seekToCallbacks       map[int]ZegoMediaPlayerSeekToCallback
}

func (mediaPlayer *mediaPlayerImpl) SetEventHandler(handler IZegoMediaPlayerEventHandler) {
	mediaPlayerLock.Lock()
	defer mediaPlayerLock.Unlock()
	mediaPlayer.eventHandler = handler
}

func (mediaPlayer *mediaPlayerImpl) SetAudioHandler(handler IZegoMediaPlayerAudioHandler) {
	mediaPlayerLock.Lock()
	defer mediaPlayerLock.Unlock()
	mediaPlayer.audioHandler = handler
}

func (mediaPlayer *mediaPlayerImpl) LoadResource(path string, callback ZegoMediaPlayerLoadResourceCallback) {
	cPath := StringToCString(path)
	defer FreeCString(cPath)
	result := C.zego_express_media_player_load_resource(cPath, C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex))
	if result != C.ZEGO_ERRCODE_COMMON_SUCCESS {
		if callback != nil {
			callback(int(result))
		}
		return
	}
	mediaPlayerLock.Lock()
	defer mediaPlayerLock.Unlock()
	mediaPlayer.loadResourceCallbacks.PushBack(callback)
}

func (mediaPlayer *mediaPlayerImpl) Start() {
	C.zego_express_media_player_start(C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex))
}

func (mediaPlayer *mediaPlayerImpl) Stop() {
	C.zego_express_media_player_stop(C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex))
}

func (mediaPlayer *mediaPlayerImpl) Pause() {
	C.zego_express_media_player_pause(C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex))
}

func (mediaPlayer *mediaPlayerImpl) Resume() {
	C.zego_express_media_player_resume(C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex))
}

func (mediaPlayer *mediaPlayerImpl) SeekTo(millisecond uint64, callback ZegoMediaPlayerSeekToCallback) {
	cSeq := C.zego_express_get_increase_seq()
	mediaPlayerLock.Lock()
	mediaPlayer.seekToCallbacks[int(cSeq)] = callback
	mediaPlayerLock.Unlock()
	C.zego_express_media_player_seek_to(C.ulonglong(millisecond), C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex), &cSeq)
}

func (mediaPlayer *mediaPlayerImpl) EnableRepeat(enable bool) {
	C.zego_express_media_player_enable_repeat(C.bool(enable), C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex))
}

func (mediaPlayer *mediaPlayerImpl) EnableAux(enable bool) {
	C.zego_express_media_player_enable_aux(C.bool(enable), C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex))
}

func (mediaPlayer *mediaPlayerImpl) SetVolume(volume int) {
	C.zego_express_media_player_set_volume(C.int(volume), C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex))
}

func (mediaPlayer *mediaPlayerImpl) SetPlayVolume(volume int) {
	C.zego_express_media_player_set_play_volume(C.int(volume), C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex))
}

func (mediaPlayer *mediaPlayerImpl) SetPublishVolume(volume int) {
	C.zego_express_media_player_set_publish_volume(C.int(volume), C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex))
}

func (mediaPlayer *mediaPlayerImpl) GetPlayVolume() int {
	var volume C.int
	C.zego_express_media_player_get_play_volume(C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex), &volume)
	return int(volume)
}

func (mediaPlayer *mediaPlayerImpl) GetPublishVolume() int {
	var volume C.int
	C.zego_express_media_player_get_publish_volume(C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex), &volume)
	return int(volume)
}

func (mediaPlayer *mediaPlayerImpl) GetIndex() int {
	return mediaPlayer.instanceIndex
}

func createEngineInner(profile ZegoEngineProfile, eventHandler IZegoEventHandler) bool {
	engineLock.Lock()
	defer engineLock.Unlock()
	if globalEngine == nil {
		globalEngine = new(engineImpl)
		result := globalEngine.init(profile, eventHandler)
		if result != ZegoErrorCodeCommonSuccess {
			globalEngine = nil
			eventHandler.OnDebugError(result, "CreateEngine", "CreateEngine failed")
			return false
		}
		return true
	}
	return false
}

func createEngine(profile ZegoEngineProfile, eventHandler IZegoEventHandler) IZegoExpressEngine {
	result := createEngineInner(profile, eventHandler)
	if result {
		for i := 0; i < maxPublishChannelCount; i++ {
			C.zego_express_enable_camera(C.bool(false), C.zego_exp_notify_device_state_mode_open, C.enum_zego_publish_channel(i))
		}
	}
	return globalEngine
}

func destroyEngine(engine IZegoExpressEngine, callback ZegoDestroyCompletionCallback) {
	engineLock.Lock()
	defer engineLock.Unlock()
	if engine != nil && engine == globalEngine {
		engineDestroyCallbackLock.Lock()
		engineDestroyCallback = callback
		engineDestroyCallbackLock.Unlock()
		C.zego_express_engine_uninit_async()
		globalEngine = nil
	}
}

func getEngine() IZegoExpressEngine {
	engineLock.RLock()
	defer engineLock.RUnlock()
	return globalEngine
}

func setEngineConfig(config ZegoEngineConfig) {
	cEngineConfig := C.struct_zego_engine_config{}
	cLogConfig := C.struct_zego_log_config{}
	if config.LogConfig != nil {
		cLogConfig.log_size = C.ulonglong(config.LogConfig.LogSize)
		cLogConfig.log_count = C.uint(config.LogConfig.LogCount)
		setCharArray(&cLogConfig.log_path[0], config.LogConfig.LogPath, C.ZEGO_EXPRESS_MAX_COMMON_LEN)
		cEngineConfig.log_config = &cLogConfig
	}
	if config.AdvancedConfig != nil {
		var advancedConfig string
		for key, value := range config.AdvancedConfig {
			advancedConfig += key + "=" + value + ";"
		}
		setCharArray(&cEngineConfig.advanced_config[0], advancedConfig, C.ZEGO_EXPRESS_MAX_SET_CONFIG_VALUE_LEN)
	}
	C.zego_express_set_engine_config(cEngineConfig)

	if value, exists := config.AdvancedConfig["max_publish_channels"]; exists {
		if count, err := strconv.Atoi(value); err == nil {
			engineLock.Lock()
			defer engineLock.Unlock()
			maxPublishChannelCount = count
		}
	}
}

func setLogConfig(config ZegoLogConfig) {
	cLogConfig := C.struct_zego_log_config{
		log_size:  C.ulonglong(config.LogSize),
		log_count: C.uint(config.LogCount),
	}
	setCharArray(&cLogConfig.log_path[0], config.LogPath, C.ZEGO_EXPRESS_MAX_COMMON_LEN)
	C.zego_express_set_log_config(cLogConfig)
}

func setRoomMode(mode ZegoRoomMode) {
	C.zego_express_set_room_mode(C.enum_zego_room_mode(mode))
}

func getVersion() string {
	var cVersion *C.char
	C.zego_express_get_version(&cVersion)
	return C.GoString(cVersion)
}

func setApiCalledCallback(callback IZegoApiCalledEventHandler) {
	callbackLock.Lock()
	defer callbackLock.Unlock()
	apiCalledCallback = callback
}

/// Utils

func setCharArray(dest *C.char, src string, maxLen C.size_t) {
	// 自动处理字符串截断和null终止
	cStr := C.CString(src)
	defer C.free(unsafe.Pointer(cStr))
	C.strncpy(dest, cStr, maxLen)
	// *(*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(dest)) + uintptr(maxLen-1))) = 0
}

// StringToCString 将Go字符串转换为C字符串(*C.char)
// 注意：调用者需负责调用C.free释放内存
func StringToCString(goStr string) *C.char {
	cStr := C.CString(goStr)
	return cStr
}

// FreeCString 释放C字符串内存
func FreeCString(cStr *C.char) {
	C.free(unsafe.Pointer(cStr))
}

func convertStream(cStream C.struct_zego_stream) ZegoStream {
	return ZegoStream{
		User:      convertUser(cStream.user),
		StreamID:  C.GoString(&cStream.stream_id[0]),
		ExtraInfo: C.GoString(&cStream.extra_info[0]),
	}
}

func convertUser(cUser C.struct_zego_user) ZegoUser {
	return ZegoUser{
		UserID:   C.GoString(&cUser.user_id[0]),
		UserName: C.GoString(&cUser.user_name[0]),
	}
}

func convertAudioFrameParam(param ZegoAudioFrameParam) C.struct_zego_audio_frame_param {
	return C.struct_zego_audio_frame_param{
		sample_rate: C.enum_zego_audio_sample_rate(param.SampleRate),
		channel:     C.enum_zego_audio_channel(param.Channel),
	}
}

func cUcharPtrToGoSlice(cData *C.uchar, cLen C.uint) []uint8 {
	len := int(cLen)
	if cData == nil || len <= 0 {
		return nil
	}
	return (*[1 << 30]uint8)(unsafe.Pointer(cData))[:len:len]
}

func goSliceToCUchar(data []uint8) (*C.uchar, C.uint) {
	var cData *C.uchar = nil
	var cLen C.uint = 0
	if len(data) > 0 {
		cData = (*C.uchar)(unsafe.Pointer(&data[0]))
		cLen = C.uint(len(data))
	}
	return cData, cLen
}

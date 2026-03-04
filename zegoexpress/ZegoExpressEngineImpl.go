package zegoexpress

/*
#cgo CFLAGS: -I${SRCDIR}/lib/include
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

// ZegoInternalPrivate.h
extern zego_handle zego_express_engine_create_handle();
extern void zego_express_engine_destroy_handle(zego_handle handle);

// 声明由Go实现的函数
extern void GoLoginResultCallback(int, char*, int, void *);
extern void GoLogoutResultCallback(int, char*, int, void *);
extern void GoOnIMSendBroadcastMessageResult(zego_error, unsigned long long msg_id, int, void *);
extern void GoOnPublisherUpdateStreamExtraInfoResult(zego_error, int, void *);
extern void GoOnPlayerAudioData(unsigned char *, unsigned int, struct zego_audio_frame_param, char *, void *);
extern void GoOnProcessRemoteAudioData(unsigned char *, unsigned int, struct zego_audio_frame_param *, char *, double, void *);
extern void GoOnDebugError(int error_code, char* func_name, char* info, void *);
extern void GoOnRoomStateUpdate(char *room_id, enum zego_room_state state, zego_error error_code, char *extend_data, void *);
extern void GoOnRoomUserUpdate(char *room_id, enum zego_update_type update_type, struct zego_user *user_list, unsigned int user_count, void *);
extern void GoOnRoomOnlineUserCountUpdate(char *room_id, int count, void *);
extern void GoOnRoomStreamUpdate(char *room_id, enum zego_update_type update_type, struct zego_stream *stream_info_list, unsigned int stream_info_count, char *extended_data, void *);
extern void GoOnRoomStreamExtraInfoUpdate(char *room_id, struct zego_stream *stream_info_list, unsigned int stream_info_count, void *);
extern void GoOnRoomStateChanged(char *room_id, enum zego_room_state_changed_reason reason, zego_error error_code, char *extended_data, void *);
extern void GoOnRoomTokenWillExpire(char *room_id, int remain_time_in_second, void *);
extern void GoOnPublisherStateUpdate(char *stream_id, enum zego_publisher_state state, zego_error error_code, char *extend_data, void *);
extern void GoOnPublisherQualityUpdate(char *stream_id, struct zego_publish_stream_quality quality, void *);
extern void GoOnPublisherStreamEvent(enum zego_stream_event event_id, char *stream_id, char *extra_info, void *);
extern void GoOnPublisherSendAudioFirstFrame(enum zego_publish_channel channel, void *);
extern void GoOnPlayerStateUpdate(char *stream_id, enum zego_player_state state, zego_error error_code, char *extend_data, void *);
extern void GoOnPlayerQualityUpdate(char *stream_id, struct zego_play_stream_quality quality, void *);
extern void GoOnPlayerRecvSei(struct zego_media_side_info info, void *);
extern void GoOnPlayerStreamEvent(enum zego_stream_event event_id, char *stream_id, char *extra_info, void *);
extern void GoOnPlayerRecvAudioFirstFrame(char *stream_id, void *);

extern void GoOnMediaPlayerStateUpdate(enum zego_media_player_state state, zego_error error_code, enum zego_media_player_instance_index instance_index, void *);
extern void GoOnMediaPlayerNetworkEvent(enum zego_media_player_network_event event, enum zego_media_player_instance_index instance_index, void *);
extern void GoOnMediaPlayerPlayingProgress(unsigned long long millisecond, enum zego_media_player_instance_index instance_index, void *);
extern void GoOnMediaPlayerRenderingProgress(unsigned long long millisecond, enum zego_media_player_instance_index instance_index, void *);
extern void GoOnMediaPlayerRecvSEI(unsigned char *data, unsigned int data_length, enum zego_media_player_instance_index instance_index, void *);
extern void GoOnMediaPlayerFirstFrameEvent(enum zego_media_player_first_frame_event event, enum zego_media_player_instance_index instance_index, void *);
extern void GoOnMediaPlayerAudioFrame(unsigned char *data, unsigned int data_length, const struct zego_audio_frame_param param, enum zego_media_player_instance_index instance_index, void *);
extern void GoOnMediaPlayerLoadFileResult(zego_error error_code, enum zego_media_player_instance_index instance_index, void *);
extern void GoOnMediaPlayerSeekTo(zego_seq seq, zego_error error_code, enum zego_media_player_instance_index instance_index, void *);
extern void GoOnEngineUninit(void *);

static void bridge_go_login_callback(zego_handle handle, zego_error code, const char *ext_data, const char *room_id, zego_seq seq, void *ctx) {
    GoLoginResultCallback(code, (char *)ext_data, seq, ctx);
}

static void bridge_go_logout_callback(zego_handle handle, zego_error code, const char *ext_data, const char *room_id, zego_seq seq, void *ctx) {
    GoLogoutResultCallback(code, (char *)ext_data, seq, ctx);
}

static void bridge_go_on_im_send_broadcast_message_result(zego_handle handle, const char *room_id, unsigned long long message_id, zego_error error_code, zego_seq seq, void *user_context) {
    GoOnIMSendBroadcastMessageResult(error_code, message_id, seq, user_context);
}

static void bridge_go_on_publisher_update_stream_extra_info_result(zego_handle handle, zego_error error_code, zego_seq seq, void *user_context) {
    GoOnPublisherUpdateStreamExtraInfoResult(error_code, seq, user_context);
}

static void bridge_go_on_player_audio_data(zego_handle handle, const unsigned char *data, unsigned int data_length, struct zego_audio_frame_param param, const char *stream_id, void *user_context) {
    GoOnPlayerAudioData((unsigned char *)data, data_length, param, (char *)stream_id, user_context);
}

static void bridge_go_on_process_remote_audio_data(zego_handle handle, unsigned char *data, unsigned int data_length, struct zego_audio_frame_param *param, const char *stream_id, double timestamp, void *user_context) {
    GoOnProcessRemoteAudioData((unsigned char *)data, data_length, param, (char *)stream_id, timestamp, user_context);
}

static void bridge_go_on_debug_error(zego_handle handle, int error_code, const char *func_name, const char *info, void *user_context) {
    GoOnDebugError(error_code, (char *)func_name, (char *)info, user_context);
}

static void bridge_go_on_room_state_update(zego_handle handle, const char *room_id, enum zego_room_state state, zego_error error_code, const char *extend_data, void *user_context) {
    GoOnRoomStateUpdate((char *)room_id, state, error_code, (char *)extend_data, user_context);
}

static void bridge_go_on_user_update(zego_handle handle, const char *room_id, enum zego_update_type update_type, const struct zego_user *user_list, unsigned int user_count, void *user_context) {
    GoOnRoomUserUpdate((char *)room_id, update_type, (struct zego_user *)user_list, user_count, user_context);
}

static void bridge_go_on_room_online_user_count_update(zego_handle handle, const char *room_id, int online_user_count, void *user_context) {
    GoOnRoomOnlineUserCountUpdate((char *)room_id, online_user_count, user_context);
}

static void bridge_go_on_room_stream_update(zego_handle handle, const char *room_id, enum zego_update_type update_type, const struct zego_stream *stream_info_list, unsigned int stream_info_count, const char *extended_data, void *user_context) {
    GoOnRoomStreamUpdate((char *)room_id, update_type, (struct zego_stream *)stream_info_list, stream_info_count, (char *)extended_data, user_context);
}

static void bridge_go_on_room_stream_extra_info_update(zego_handle handle, const char *room_id, const struct zego_stream *stream_info_list, unsigned int stream_info_count, void *user_context) {
    GoOnRoomStreamExtraInfoUpdate((char *)room_id, (struct zego_stream *)stream_info_list, stream_info_count, user_context);
}

static void bridge_go_on_room_state_changed(zego_handle handle, const char *room_id, enum zego_room_state_changed_reason reason, zego_error error_code, const char *extended_data, void *user_context) {
    GoOnRoomStateChanged((char *)room_id, reason, error_code, (char *)extended_data, user_context);
}

static void bridge_go_on_room_token_will_expire(zego_handle handle, const char *room_id, int remain_time_in_second, void *user_context) {
    GoOnRoomTokenWillExpire((char *)room_id, remain_time_in_second, user_context);
}

static void bridge_go_on_publisher_state_update(zego_handle handle, const char *stream_id, enum zego_publisher_state state, zego_error error_code, const char *extend_data, void *user_context) {
    GoOnPublisherStateUpdate((char *)stream_id, state, error_code, (char *)extend_data, user_context);
}

static void bridge_go_on_publisher_quality_update(zego_handle handle, const char *stream_id, struct zego_publish_stream_quality quality, void *user_context) {
    GoOnPublisherQualityUpdate((char *)stream_id, quality, user_context);
}

static void bridge_go_on_publisher_stream_event(zego_handle handle, enum zego_stream_event event_id, const char *stream_id, const char *extra_info, void *user_context) {
    GoOnPublisherStreamEvent(event_id, (char *)stream_id, (char *)extra_info, user_context);
}

static void bridge_go_on_publisher_send_audio_first_frame(zego_handle handle, enum zego_publish_channel channel, void *user_context) {
    GoOnPublisherSendAudioFirstFrame(channel, user_context);
}

static void bridge_go_on_player_state_update(zego_handle handle, const char *stream_id, enum zego_player_state state, zego_error error_code, const char *extend_data, void *user_context) {
    GoOnPlayerStateUpdate((char *)stream_id, state, error_code, (char *)extend_data, user_context);
}

static void bridge_go_on_player_quality_update(zego_handle handle, const char *stream_id, struct zego_play_stream_quality quality, void *user_context) {
    GoOnPlayerQualityUpdate((char *)stream_id, quality, user_context);
}

static void bridge_go_on_player_recv_sei(zego_handle handle, struct zego_media_side_info info, void *user_context) {
    GoOnPlayerRecvSei(info, user_context);
}

static void bridge_go_on_player_stream_event(zego_handle handle, enum zego_stream_event event_id, const char *stream_id, const char *extra_info, void *user_context) {
    GoOnPlayerStreamEvent(event_id, (char *)stream_id, (char *)extra_info, user_context);
}

static void bridge_go_on_player_recv_audio_first_frame(zego_handle handle, const char *stream_id, void *user_context) {
    GoOnPlayerRecvAudioFirstFrame((char *)stream_id, user_context);
}

static void bridge_go_on_media_player_state_update(zego_handle handle, enum zego_media_player_state state, zego_error error_code, enum zego_media_player_instance_index instance_index, void *user_context) {
    GoOnMediaPlayerStateUpdate(state, error_code, instance_index, user_context);
}

static void bridge_go_on_media_player_network_event(zego_handle handle, enum zego_media_player_network_event event, enum zego_media_player_instance_index instance_index, void *user_context) {
    GoOnMediaPlayerNetworkEvent(event, instance_index, user_context);
}

static void bridge_go_on_media_player_playing_progress(zego_handle handle, unsigned long long millisecond, enum zego_media_player_instance_index instance_index, void *user_context) {
    GoOnMediaPlayerPlayingProgress(millisecond, instance_index, user_context);
}

static void bridge_go_on_media_player_rendering_progress(zego_handle handle, unsigned long long millisecond, enum zego_media_player_instance_index instance_index, void *user_context) {
    GoOnMediaPlayerRenderingProgress(millisecond, instance_index, user_context);
}

static void bridge_go_on_media_player_recv_sei(zego_handle handle, const unsigned char *data, unsigned int data_length, enum zego_media_player_instance_index instance_index, void *user_context) {
    GoOnMediaPlayerRecvSEI((char *)data, data_length, instance_index, user_context);
}

static void bridge_go_on_media_player_first_frame_event(zego_handle handle, enum zego_media_player_first_frame_event event, enum zego_media_player_instance_index instance_index, void *user_context) {
    GoOnMediaPlayerFirstFrameEvent(event, instance_index, user_context);
}

static void bridge_go_on_media_player_audio_frame(zego_handle handle, unsigned char *data, unsigned int data_length, const struct zego_audio_frame_param param, enum zego_media_player_instance_index instance_index, void *user_context) {
    GoOnMediaPlayerAudioFrame((char *)data, data_length, param, instance_index, user_context);
}

static void bridge_go_on_mediaplayer_load_file_result(zego_handle handle, zego_error error_code, enum zego_media_player_instance_index instance_index, void *user_context) {
    GoOnMediaPlayerLoadFileResult(error_code, instance_index, user_context);
}

static void bridge_go_on_media_player_seek_to(zego_handle handle, zego_seq seq, zego_error error_code, enum zego_media_player_instance_index instance_index, void *user_context) {
    GoOnMediaPlayerSeekTo(seq, error_code, instance_index, user_context);
}

static void bridge_go_on_engine_uninit(zego_handle handle, void *user_context) {
    GoOnEngineUninit(user_context);
}

static void zego_express_go_bridge_init(zego_handle handle, void *user_context) {
    zego_register_room_login_result_callback(handle, bridge_go_login_callback, user_context);
    zego_register_room_logout_result_callback(handle, bridge_go_logout_callback, user_context);
    zego_register_im_send_broadcast_message_result_callback(handle, bridge_go_on_im_send_broadcast_message_result, user_context);
    zego_register_publisher_update_stream_extra_info_result_callback(handle, bridge_go_on_publisher_update_stream_extra_info_result, user_context);
    zego_register_player_audio_data_callback(handle, bridge_go_on_player_audio_data, user_context);
    zego_register_process_remote_audio_data_callback(handle, bridge_go_on_process_remote_audio_data, user_context);
    zego_register_debug_error_callback(handle, bridge_go_on_debug_error, user_context);
    zego_register_room_state_update_callback(handle, bridge_go_on_room_state_update, user_context);
    zego_register_room_user_update_callback(handle, bridge_go_on_user_update, user_context);
    zego_register_room_online_user_count_update_callback(handle, bridge_go_on_room_online_user_count_update, user_context);
    zego_register_room_stream_update_callback(handle, bridge_go_on_room_stream_update, user_context);
    zego_register_room_stream_extra_info_update_callback(handle, bridge_go_on_room_stream_extra_info_update, user_context);
    zego_register_room_state_changed_callback(handle, bridge_go_on_room_state_changed, user_context);
    zego_register_room_token_will_expire_callback(handle, bridge_go_on_room_token_will_expire, user_context);
    zego_register_publisher_state_update_callback(handle, bridge_go_on_publisher_state_update, user_context);
    zego_register_publisher_quality_update_callback(handle, bridge_go_on_publisher_quality_update, user_context);
    zego_register_publisher_stream_event_callback(handle, bridge_go_on_publisher_stream_event, user_context);
    zego_register_publisher_send_audio_first_frame_callback(handle, bridge_go_on_publisher_send_audio_first_frame, user_context);
    zego_register_player_state_update_callback(handle, bridge_go_on_player_state_update, user_context);
    zego_register_player_quality_update_callback(handle, bridge_go_on_player_quality_update, user_context);
    zego_register_player_recv_media_side_info_callback(handle, bridge_go_on_player_recv_sei, user_context);
    zego_register_player_stream_event_callback(handle, bridge_go_on_player_stream_event, user_context);
    zego_register_player_recv_audio_first_frame_callback(handle, bridge_go_on_player_recv_audio_first_frame, user_context);
    zego_register_media_player_state_update_callback(handle, bridge_go_on_media_player_state_update, user_context);
    zego_register_media_player_network_event_callback(handle, bridge_go_on_media_player_network_event, user_context);
    zego_register_media_player_playing_progress_callback(handle, bridge_go_on_media_player_playing_progress, user_context);
    zego_register_media_player_rendering_progress_callback(handle, bridge_go_on_media_player_rendering_progress, user_context);
    zego_register_media_player_recv_sei_callback(handle, bridge_go_on_media_player_recv_sei, user_context);
    zego_register_media_player_first_frame_event_callback(handle, bridge_go_on_media_player_first_frame_event, user_context);
    zego_register_media_player_audio_frame_callback(handle, bridge_go_on_media_player_audio_frame, user_context);
    zego_register_media_player_load_resource_callback(handle, bridge_go_on_mediaplayer_load_file_result, user_context);
    zego_register_media_player_seek_to_callback(handle, bridge_go_on_media_player_seek_to, user_context);
    zego_register_engine_uninit_callback(handle, bridge_go_on_engine_uninit, user_context);
}

*/
import "C"
import (
	"container/list"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"unsafe"
)

func init() {
	fmt.Println("zegoexpress init")

	gCallbackHandler = &callbackHandler{
		callbackChan: make(chan func(), 131072), // 1024 * 1024 / 8 ~ 1MB
	}
	go gCallbackHandler.processLoop()
}

var (
	engineDestroyCallbackLock sync.Mutex
	engineDestroyCallbacks    map[unsafe.Pointer]ZegoDestroyCompletionCallback

	callbackEventLock    sync.Mutex
	callbackEventHandler IZegoCallbackEventHandler

	gCallbackHandler *callbackHandler
)

//export GoLoginResultCallback
func GoLoginResultCallback(errorCode C.zego_error, extendedData *C.char, seq C.zego_seq, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	goExtendedData := ""
	if extendedData != nil {
		goExtendedData = C.GoString(extendedData)
	}
	callbackFunc := func() {
		engine.callbackLock.Lock()
		defer engine.callbackLock.Unlock()

		callback, ok := engine.roomLoginCallback[int(seq)]
		if !ok {
			return
		}

		if callback != nil {
			callback(int(errorCode), goExtendedData)
		}

		delete(engine.roomLoginCallback, int(seq))
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoLogoutResultCallback
func GoLogoutResultCallback(errorCode C.zego_error, extendedData *C.char, seq C.zego_seq, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	goExtendedData := ""
	if extendedData != nil {
		goExtendedData = C.GoString(extendedData)
	}
	callbackFunc := func() {
		engine.callbackLock.Lock()
		defer engine.callbackLock.Unlock()

		callback, ok := engine.roomLogoutCallback[int(seq)]
		if !ok {
			return
		}

		if callback != nil {
			callback(int(errorCode), goExtendedData)
		}

		delete(engine.roomLogoutCallback, int(seq))
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnIMSendBroadcastMessageResult
func GoOnIMSendBroadcastMessageResult(errorCode C.zego_error, messageID C.ulonglong, seq C.zego_seq, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	callbackFunc := func() {
		engine.callbackLock.Lock()
		defer engine.callbackLock.Unlock()

		callback, ok := engine.imSendBroadcastMessageCallback[int(seq)]
		if !ok {
			return
		}

		if callback != nil {
			callback(int(errorCode), uint64(messageID))
		}

		delete(engine.imSendBroadcastMessageCallback, int(seq))
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnPublisherUpdateStreamExtraInfoResult
func GoOnPublisherUpdateStreamExtraInfoResult(errorCode C.zego_error, seq C.zego_seq, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	callbackFunc := func() {
		engine.callbackLock.Lock()
		defer engine.callbackLock.Unlock()

		callback, ok := engine.setStreamExtraInfoCallback[int(seq)]
		if !ok {
			return
		}

		if callback != nil {
			callback(int(errorCode))
		}

		delete(engine.setStreamExtraInfoCallback, int(seq))
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnPlayerAudioData
func GoOnPlayerAudioData(data *C.uchar, dataLen C.uint, param C.struct_zego_audio_frame_param, streamID *C.char, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	engine.handlerLock.RLock()
	defer engine.handlerLock.RUnlock()
	handler := engine.audioDataHandler
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
func GoOnProcessRemoteAudioData(data *C.uchar, dataLen C.uint, param *C.struct_zego_audio_frame_param, streamID *C.char, timestamp float64, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	engine.handlerLock.RLock()
	defer engine.handlerLock.RUnlock()
	handler := engine.customAudioProcessHandler
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
func GoOnDebugError(errorCode C.int, funcName *C.char, info *C.char, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	engine.handlerLock.RLock()
	defer engine.handlerLock.RUnlock()
	handler := engine.eventHandler
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
func GoOnRoomStateUpdate(roomID *C.char, state C.enum_zego_room_state, errorCode C.zego_error, data *C.char, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	goRoomID := C.GoString(roomID)
	goData := ""
	if data != nil {
		goData = C.GoString(data)
	}
	callbackFunc := func() {
		engine.handlerLock.RLock()
		defer engine.handlerLock.RUnlock()
		handler := engine.eventHandler
		if handler == nil {
			return
		}
		handler.OnRoomStateUpdate(goRoomID, ZegoRoomState(state), int(errorCode), goData)
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnRoomUserUpdate
func GoOnRoomUserUpdate(roomID *C.char, updateType C.enum_zego_update_type, userList *C.struct_zego_user, userCount C.uint, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	goRoomID := C.GoString(roomID)
	goUserList := make([]ZegoUser, 0)
	if userList != nil && userCount > 0 {
		cUsers := unsafe.Slice(userList, userCount)

		for i := 0; i < int(userCount); i++ {
			user := cUsers[i]
			goUserList = append(goUserList, convertUser(user))
		}
	}
	callbackFunc := func() {
		engine.handlerLock.RLock()
		defer engine.handlerLock.RUnlock()
		handler := engine.eventHandler
		if handler == nil {
			return
		}
		handler.OnRoomUserUpdate(goRoomID, ZegoUpdateType(updateType), goUserList)
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnRoomOnlineUserCountUpdate
func GoOnRoomOnlineUserCountUpdate(roomID *C.char, count C.int, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	goRoomID := C.GoString(roomID)
	callbackFunc := func() {
		engine.handlerLock.RLock()
		defer engine.handlerLock.RUnlock()
		handler := engine.eventHandler
		if handler == nil {
			return
		}
		handler.OnRoomOnlineUserCountUpdate(goRoomID, int(count))
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnRoomStreamUpdate
func GoOnRoomStreamUpdate(roomID *C.char, updateType C.enum_zego_update_type, streamInfoList *C.struct_zego_stream, streamInfoCount C.uint, data *C.char, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	goRoomID := C.GoString(roomID)
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
	callbackFunc := func() {
		engine.handlerLock.RLock()
		defer engine.handlerLock.RUnlock()
		handler := engine.eventHandler
		if handler == nil {
			return
		}
		handler.OnRoomStreamUpdate(goRoomID, ZegoUpdateType(updateType), streamList, goData)
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnRoomStreamExtraInfoUpdate
func GoOnRoomStreamExtraInfoUpdate(roomID *C.char, streamInfoList *C.struct_zego_stream, streamInfoCount C.uint, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	goRoomID := C.GoString(roomID)
	streamList := make([]ZegoStream, 0)
	if streamInfoList != nil && streamInfoCount > 0 {
		cStreams := unsafe.Slice(streamInfoList, streamInfoCount)

		for i := 0; i < int(streamInfoCount); i++ {
			stream := cStreams[i]
			streamList = append(streamList, convertStream(stream))
		}
	}
	callbackFunc := func() {
		engine.handlerLock.RLock()
		defer engine.handlerLock.RUnlock()
		handler := engine.eventHandler
		if handler == nil {
			return
		}
		handler.OnRoomStreamExtraInfoUpdate(goRoomID, streamList)
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnRoomStateChanged
func GoOnRoomStateChanged(roomID *C.char, reason C.enum_zego_room_state_changed_reason, errorCode C.zego_error, data *C.char, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	goRoomID := C.GoString(roomID)
	goData := ""
	if data != nil {
		goData = C.GoString(data)
	}
	callbackFunc := func() {
		engine.handlerLock.RLock()
		defer engine.handlerLock.RUnlock()
		handler := engine.eventHandler
		if handler == nil {
			return
		}
		handler.OnRoomStateChanged(goRoomID, ZegoRoomStateChangedReason(reason), int(errorCode), goData)
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnRoomTokenWillExpire
func GoOnRoomTokenWillExpire(roomID *C.char, remainTimeInSecond C.int, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	goRoomID := C.GoString(roomID)
	callbackFunc := func() {
		engine.handlerLock.RLock()
		defer engine.handlerLock.RUnlock()
		handler := engine.eventHandler
		if handler == nil {
			return
		}
		handler.OnRoomTokenWillExpire(goRoomID, int(remainTimeInSecond))
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnPublisherStateUpdate
func GoOnPublisherStateUpdate(streamID *C.char, state C.enum_zego_publisher_state, errorCode C.zego_error, data *C.char, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	goStreamID := C.GoString(streamID)
	goData := ""
	if data != nil {
		goData = C.GoString(data)
	}
	callbackFunc := func() {
		engine.handlerLock.RLock()
		defer engine.handlerLock.RUnlock()
		handler := engine.eventHandler
		if handler == nil {
			return
		}
		handler.OnPublisherStateUpdate(goStreamID, ZegoPublisherState(state), int(errorCode), goData)
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnPublisherQualityUpdate
func GoOnPublisherQualityUpdate(streamID *C.char, quality C.struct_zego_publish_stream_quality, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	goStreamID := C.GoString(streamID)
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
	callbackFunc := func() {
		engine.handlerLock.RLock()
		defer engine.handlerLock.RUnlock()
		handler := engine.eventHandler
		if handler == nil {
			return
		}
		handler.OnPublisherQualityUpdate(goStreamID, goQuality)
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnPublisherStreamEvent
func GoOnPublisherStreamEvent(eventID C.enum_zego_stream_event, streamID *C.char, data *C.char, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	goStreamID := C.GoString(streamID)
	goData := ""
	if data != nil {
		goData = C.GoString(data)
	}
	callbackFunc := func() {
		engine.handlerLock.RLock()
		defer engine.handlerLock.RUnlock()
		handler := engine.eventHandler
		if handler == nil {
			return
		}
		handler.OnPublisherStreamEvent(ZegoStreamEvent(eventID), goStreamID, goData)
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnPublisherSendAudioFirstFrame
func GoOnPublisherSendAudioFirstFrame(channel C.enum_zego_publish_channel, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	callbackFunc := func() {
		engine.handlerLock.RLock()
		defer engine.handlerLock.RUnlock()
		handler := engine.eventHandler
		if handler == nil {
			return
		}
		handler.OnPublisherSendAudioFirstFrame(ZegoPublishChannel(channel))
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnPlayerStateUpdate
func GoOnPlayerStateUpdate(streamID *C.char, state C.enum_zego_player_state, errorCode C.zego_error, data *C.char, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	goStreamID := C.GoString(streamID)
	goData := ""
	if data != nil {
		goData = C.GoString(data)
	}
	callbackFunc := func() {
		engine.handlerLock.RLock()
		defer engine.handlerLock.RUnlock()
		handler := engine.eventHandler
		if handler == nil {
			return
		}
		handler.OnPlayerStateUpdate(goStreamID, ZegoPlayerState(state), int(errorCode), goData)
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnPlayerQualityUpdate
func GoOnPlayerQualityUpdate(streamID *C.char, quality C.struct_zego_play_stream_quality, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	goStreamID := C.GoString(streamID)
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
	callbackFunc := func() {
		engine.handlerLock.RLock()
		defer engine.handlerLock.RUnlock()
		handler := engine.eventHandler
		if handler == nil {
			return
		}
		handler.OnPlayerQualityUpdate(goStreamID, goQuality)
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnPlayerRecvSei
func GoOnPlayerRecvSei(info C.struct_zego_media_side_info, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	engine.handlerLock.RLock()
	defer engine.handlerLock.RUnlock()
	handler := engine.eventHandler
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
func GoOnPlayerStreamEvent(eventID C.enum_zego_stream_event, streamID *C.char, data *C.char, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	goStreamID := C.GoString(streamID)
	goData := ""
	if data != nil {
		goData = C.GoString(data)
	}
	callbackFunc := func() {
		engine.handlerLock.RLock()
		defer engine.handlerLock.RUnlock()
		handler := engine.eventHandler
		if handler == nil {
			return
		}
		handler.OnPlayerStreamEvent(ZegoStreamEvent(eventID), goStreamID, goData)
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnPlayerRecvAudioFirstFrame
func GoOnPlayerRecvAudioFirstFrame(streamID *C.char, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	goStreamID := C.GoString(streamID)
	callbackFunc := func() {
		engine.handlerLock.RLock()
		defer engine.handlerLock.RUnlock()
		handler := engine.eventHandler
		if handler == nil {
			return
		}
		handler.OnPlayerRecvAudioFirstFrame(goStreamID)
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnMediaPlayerStateUpdate
func GoOnMediaPlayerStateUpdate(state C.enum_zego_media_player_state, errorCode C.zego_error, index C.enum_zego_media_player_instance_index, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	callbackFunc := func() {
		engine.mediaPlayerLock.Lock()
		defer engine.mediaPlayerLock.Unlock()
		if mediaPlayer, ok := engine.mediaPlayerImplMap[int(index)]; ok {
			mediaPlayer.handlerLock.Lock()
			defer mediaPlayer.handlerLock.Unlock()
			handler := mediaPlayer.eventHandler
			if handler != nil {
				handler.OnMediaPlayerStateUpdate(mediaPlayer, ZegoMediaPlayerState(state), int(errorCode))
			}
		}
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnMediaPlayerNetworkEvent
func GoOnMediaPlayerNetworkEvent(event C.enum_zego_media_player_network_event, index C.enum_zego_media_player_instance_index, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	callbackFunc := func() {
		engine.mediaPlayerLock.Lock()
		defer engine.mediaPlayerLock.Unlock()
		if mediaPlayer, ok := engine.mediaPlayerImplMap[int(index)]; ok {
			mediaPlayer.handlerLock.Lock()
			defer mediaPlayer.handlerLock.Unlock()
			handler := mediaPlayer.eventHandler
			if handler != nil {
				handler.OnMediaPlayerNetworkEvent(mediaPlayer, ZegoMediaPlayerNetworkEvent(event))
			}
		}
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnMediaPlayerPlayingProgress
func GoOnMediaPlayerPlayingProgress(millisecond C.ulonglong, index C.enum_zego_media_player_instance_index, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	engine.mediaPlayerLock.Lock()
	defer engine.mediaPlayerLock.Unlock()
	if mediaPlayer, ok := engine.mediaPlayerImplMap[int(index)]; ok {
		mediaPlayer.handlerLock.Lock()
		defer mediaPlayer.handlerLock.Unlock()
		handler := mediaPlayer.eventHandler
		if handler != nil {
			handler.OnMediaPlayerPlayingProgress(mediaPlayer, uint64(millisecond))
		}
	}
}

//export GoOnMediaPlayerRenderingProgress
func GoOnMediaPlayerRenderingProgress(millisecond C.ulonglong, index C.enum_zego_media_player_instance_index, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	engine.mediaPlayerLock.Lock()
	defer engine.mediaPlayerLock.Unlock()
	if mediaPlayer, ok := engine.mediaPlayerImplMap[int(index)]; ok {
		mediaPlayer.handlerLock.Lock()
		defer mediaPlayer.handlerLock.Unlock()
		handler := mediaPlayer.eventHandler
		if handler != nil {
			handler.OnMediaPlayerRenderingProgress(mediaPlayer, uint64(millisecond))
		}
	}
}

//export GoOnMediaPlayerRecvSEI
func GoOnMediaPlayerRecvSEI(data *C.uchar, dataLen C.uint, index C.enum_zego_media_player_instance_index, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	engine.mediaPlayerLock.Lock()
	defer engine.mediaPlayerLock.Unlock()
	if mediaPlayer, ok := engine.mediaPlayerImplMap[int(index)]; ok {
		mediaPlayer.handlerLock.Lock()
		defer mediaPlayer.handlerLock.Unlock()
		handler := mediaPlayer.eventHandler
		if handler != nil {
			goData := cUcharPtrToGoSlice(data, dataLen)
			handler.OnMediaPlayerRecvSEI(mediaPlayer, goData)
		}
	}
}

//export GoOnMediaPlayerFirstFrameEvent
func GoOnMediaPlayerFirstFrameEvent(event C.enum_zego_media_player_first_frame_event, index C.enum_zego_media_player_instance_index, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	callbackFunc := func() {
		engine.mediaPlayerLock.Lock()
		defer engine.mediaPlayerLock.Unlock()
		if mediaPlayer, ok := engine.mediaPlayerImplMap[int(index)]; ok {
			mediaPlayer.handlerLock.Lock()
			defer mediaPlayer.handlerLock.Unlock()
			handler := mediaPlayer.eventHandler
			if handler != nil {
				handler.OnMediaPlayerFirstFrameEvent(mediaPlayer, ZegoMediaPlayerFirstFrameEvent(event))
			}
		}
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnMediaPlayerAudioFrame
func GoOnMediaPlayerAudioFrame(data *C.uchar, dataLen C.uint, param C.struct_zego_audio_frame_param, index C.enum_zego_media_player_instance_index, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	engine.mediaPlayerLock.Lock()
	defer engine.mediaPlayerLock.Unlock()
	if mediaPlayer, ok := engine.mediaPlayerImplMap[int(index)]; ok {
		mediaPlayer.handlerLock.Lock()
		defer mediaPlayer.handlerLock.Unlock()
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
func GoOnMediaPlayerLoadFileResult(errorCode C.zego_error, index C.enum_zego_media_player_instance_index, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	callbackFunc := func() {
		engine.mediaPlayerLock.Lock()
		defer engine.mediaPlayerLock.Unlock()
		if mediaPlayer, ok := engine.mediaPlayerImplMap[int(index)]; ok {
			mediaPlayer.handlerLock.Lock()
			defer mediaPlayer.handlerLock.Unlock()
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
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnMediaPlayerSeekTo
func GoOnMediaPlayerSeekTo(seq C.zego_seq, errorCode C.zego_error, index C.enum_zego_media_player_instance_index, ctx unsafe.Pointer) {
	engine := (*engineImpl)(ctx)
	callbackFunc := func() {
		engine.mediaPlayerLock.Lock()
		defer engine.mediaPlayerLock.Unlock()
		if mediaPlayer, ok := engine.mediaPlayerImplMap[int(index)]; ok {
			mediaPlayer.handlerLock.Lock()
			defer mediaPlayer.handlerLock.Unlock()
			callbacks := mediaPlayer.seekToCallbacks
			if callback, ok := callbacks[int(seq)]; ok {
				if callback != nil {
					callback(int(errorCode))
				}
				delete(callbacks, int(seq))
			}
		}
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

//export GoOnEngineUninit
func GoOnEngineUninit(ctx unsafe.Pointer) {
	callbackFunc := func() {
		engineDestroyCallbackLock.Lock()
		defer engineDestroyCallbackLock.Unlock()
		if callback, ok := engineDestroyCallbacks[ctx]; ok {
			if callback != nil {
				callback()
			}
			delete(engineDestroyCallbacks, ctx)
		}
	}
	gCallbackHandler.dispatchInCallbackGoroutine(callbackFunc)
}

type engineImpl struct {
	handle C.zego_handle
	pinner runtime.Pinner

	handlerLock               sync.RWMutex
	eventHandler              IZegoEventHandler
	audioDataHandler          IZegoAudioDataHandler
	customAudioProcessHandler IZegoCustomAudioProcessHandler

	callbackLock                   sync.Mutex
	roomLoginCallback              map[int]ZegoRoomLoginCallback
	roomLogoutCallback             map[int]ZegoRoomLogoutCallback
	imSendBroadcastMessageCallback map[int]ZegoIMSendBroadcastMessageCallback
	setStreamExtraInfoCallback     map[int]ZegoPublisherSetStreamExtraInfoCallback

	mediaPlayerLock    sync.Mutex
	mediaPlayerImplMap map[int]*mediaPlayerImpl
}

func NewEngineImpl() *engineImpl {
	result := &engineImpl{
		roomLoginCallback:              make(map[int]ZegoRoomLoginCallback),
		roomLogoutCallback:             make(map[int]ZegoRoomLogoutCallback),
		imSendBroadcastMessageCallback: make(map[int]ZegoIMSendBroadcastMessageCallback),
		setStreamExtraInfoCallback:     make(map[int]ZegoPublisherSetStreamExtraInfoCallback),
		mediaPlayerImplMap:             make(map[int]*mediaPlayerImpl),
	}
	result.pinner.Pin(result)
	return result
}

func (e *engineImpl) init(profile ZegoMultiEngineProfile, handler IZegoEventHandler) int {
	e.handlerLock.Lock()
	e.eventHandler = handler
	e.handlerLock.Unlock()

	e.handle = C.zego_express_engine_create_handle()
	C.zego_express_go_bridge_init(e.handle, unsafe.Pointer(e))
	C.zego_express_set_room_mode(e.handle, C.enum_zego_room_mode(profile.RoomMode))

	if profile.AdvancedConfig == nil {
		profile.AdvancedConfig = make(map[string]string)
	}
	profile.AdvancedConfig["thirdparty_framework_info"] = "golang"
	e.SetEngineConfig(profile.AdvancedConfig)

	cProfile := C.struct_zego_engine_profile{
		app_id:   C.uint(profile.AppID),
		scenario: C.enum_zego_scenario(profile.Scenario),
	}
	setCharArray(&cProfile.app_sign[0], profile.AppSign, C.ZEGO_EXPRESS_MAX_APPSIGN_LEN)
	return int(C.zego_express_engine_init_with_profile(e.handle, cProfile))
}

func (e *engineImpl) SetEngineConfig(advancedConfig map[string]string) {
	if advancedConfig == nil {
		return
	}
	cEngineConfig := C.struct_zego_engine_config{}
	var advancedStr string
	for key, value := range advancedConfig {
		advancedStr += key + "=" + value + ";"
	}
	setCharArray(&cEngineConfig.advanced_config[0], advancedStr, C.ZEGO_EXPRESS_MAX_SET_CONFIG_VALUE_LEN)
	C.zego_express_set_engine_config_with_instance(e.handle, cEngineConfig)
}

func (e *engineImpl) EnableDebugAssistant(enable bool) {
	C.zego_express_enable_debug_assistant(e.handle, C.bool(enable))
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
	C.zego_express_login_room_with_callback(e.handle, cRoomID, cZegoUser, cRoomConfigPtr, &seq)
	if callback != nil {
		e.callbackLock.Lock()
		defer e.callbackLock.Unlock()
		e.roomLoginCallback[int(seq)] = callback
	}
}

func (e *engineImpl) LogoutRoom(roomID string, callback ZegoRoomLogoutCallback) {
	var seq C.int
	cRoomID := StringToCString(roomID)
	defer FreeCString(cRoomID)
	C.zego_express_logout_room_with_callback(e.handle, cRoomID, &seq)
	if callback != nil {
		e.callbackLock.Lock()
		defer e.callbackLock.Unlock()
		e.roomLogoutCallback[int(seq)] = callback
	}
}

func (e *engineImpl) RenewToken(roomID string, token string) {
	cRoomID := StringToCString(roomID)
	defer FreeCString(cRoomID)
	cToken := StringToCString(token)
	defer FreeCString(cToken)
	C.zego_express_renew_token(e.handle, cRoomID, cToken)
}

func (e *engineImpl) SendBroadcastMessage(roomID string, message string, callback ZegoIMSendBroadcastMessageCallback) {
	var seq C.int

	cRoomID := StringToCString(roomID)
	defer FreeCString(cRoomID)
	cMessage := StringToCString(message)
	defer FreeCString(cMessage)

	C.zego_express_send_broadcast_message(e.handle, cRoomID, cMessage, &seq)
	if callback != nil {
		e.callbackLock.Lock()
		defer e.callbackLock.Unlock()
		e.imSendBroadcastMessageCallback[int(seq)] = callback
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
	C.zego_express_get_room_stream_list(e.handle, cRoomID, C.enum_zego_room_stream_list_type(streamListType), &cResult)
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
	C.zego_express_free_room_stream_list(e.handle, cResult)
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

	C.zego_express_start_publishing_stream_with_config(e.handle, cStreamID, cConfig, C.enum_zego_publish_channel(channel))
}

func (e *engineImpl) StopPublishingStream(channel ZegoPublishChannel) {
	C.zego_express_stop_publishing_stream(e.handle, C.enum_zego_publish_channel(channel))
}

func (e *engineImpl) SetStreamExtraInfo(extraInfo string, callback ZegoPublisherSetStreamExtraInfoCallback, channel ZegoPublishChannel) {
	var seq C.int

	cExtraInfo := StringToCString(extraInfo)
	defer FreeCString(cExtraInfo)

	C.zego_express_set_stream_extra_info(e.handle, cExtraInfo, C.enum_zego_publish_channel(channel), &seq)
	if callback != nil {
		e.callbackLock.Lock()
		defer e.callbackLock.Unlock()
		e.setStreamExtraInfoCallback[int(seq)] = callback
	}
}

func (e *engineImpl) SetAudioConfig(config ZegoAudioConfig, channel ZegoPublishChannel) {
	cConfig := C.struct_zego_audio_config{
		bitrate:  C.int(config.Bitrate),
		channel:  C.enum_zego_audio_channel(config.Channel),
		codec_id: C.enum_zego_audio_codec_id(config.CodecID),
	}
	C.zego_express_set_audio_config_by_channel(e.handle, cConfig, C.enum_zego_publish_channel(channel))
}

func (e *engineImpl) EnableAEC(enable bool) {
	C.zego_express_enable_aec(e.handle, C.bool(enable))
}

func (e *engineImpl) EnableAGC(enable bool) {
	C.zego_express_enable_agc(e.handle, C.bool(enable))
}

func (e *engineImpl) EnableANS(enable bool) {
	C.zego_express_enable_ans(e.handle, C.bool(enable))
}

func (e *engineImpl) EnableCustomAudioIO(enable bool, config *ZegoCustomAudioConfig, channel ZegoPublishChannel) {
	cConfig := C.struct_zego_custom_audio_config{
		source_type: C.zego_audio_source_type_default,
	}
	if config != nil {
		cConfig.source_type = C.enum_zego_audio_source_type(config.SourceType)
	}
	C.zego_express_enable_custom_audio_io(e.handle, C.bool(enable), &cConfig, C.enum_zego_publish_channel(channel))
}

func (e *engineImpl) SendSEI(data []uint8, channel ZegoPublishChannel) {
	cData, cLen := goSliceToCUchar(data)
	C.zego_express_send_sei(e.handle, cData, cLen, C.enum_zego_publish_channel(channel))
}

func (e *engineImpl) SetAudioDataHandler(handler IZegoAudioDataHandler) {
	e.handlerLock.Lock()
	defer e.handlerLock.Unlock()
	e.audioDataHandler = handler
}

func (e *engineImpl) StartAudioDataObserver(observerBitMask uint32, param ZegoAudioFrameParam) {
	C.zego_express_start_audio_data_observer(e.handle, C.uint(observerBitMask), convertAudioFrameParam(param))
}

func (e *engineImpl) StopAudioDataObserver() {
	C.zego_express_stop_audio_data_observer(e.handle)
}

func (e *engineImpl) StartPlayingStream(streamID string, config *ZegoPlayerConfig) {
	cStreamID := StringToCString(streamID)
	defer FreeCString(cStreamID)
	if config == nil {
		C.zego_express_start_playing_stream(e.handle, cStreamID, nil)
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
	C.zego_express_start_playing_stream_with_config(e.handle, cStreamID, nil, cConfig)
}

func (e *engineImpl) StopPlayingStream(streamID string) {
	cStreamID := StringToCString(streamID)
	defer FreeCString(cStreamID)
	C.zego_express_stop_playing_stream(e.handle, cStreamID)
}

func (e *engineImpl) SendCustomAudioCapturePCMData(data []uint8, param ZegoAudioFrameParam, channel ZegoPublishChannel) {
	cData, cLen := goSliceToCUchar(data)
	cParam := convertAudioFrameParam(param)
	C.zego_express_send_custom_audio_capture_pcm_data(e.handle, cData, cLen, cParam, C.enum_zego_publish_channel(channel))
}

func (e *engineImpl) FetchCustomAudioRenderPCMData(data []uint8, param ZegoAudioFrameParam) {
	cData, cLen := goSliceToCUchar(data)
	cParam := convertAudioFrameParam(param)
	C.zego_express_fetch_custom_audio_render_pcm_data(e.handle, cData, cLen, cParam)
}

func (e *engineImpl) SetCustomAudioProcessHandler(handle IZegoCustomAudioProcessHandler) {
	e.handlerLock.Lock()
	defer e.handlerLock.Unlock()
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
	C.zego_express_enable_custom_audio_remote_processing(e.handle, C.bool(enable), cConfigPtr)
}

func (e *engineImpl) CreateMediaPlayer() IZegoMediaPlayer {
	var index C.enum_zego_media_player_instance_index = C.zego_media_player_instance_index_null
	C.zego_express_create_media_player(e.handle, &index)
	if index == C.zego_media_player_instance_index_null {
		return nil
	}
	e.mediaPlayerLock.Lock()
	defer e.mediaPlayerLock.Unlock()
	mediaPlayer := new(mediaPlayerImpl)
	mediaPlayer.handle = e.handle
	mediaPlayer.instanceIndex = int(index)
	mediaPlayer.loadResourceCallbacks = list.New()
	mediaPlayer.seekToCallbacks = make(map[int]ZegoMediaPlayerSeekToCallback)
	e.mediaPlayerImplMap[int(index)] = mediaPlayer
	return mediaPlayer
}

func (e *engineImpl) DestroyMediaPlayer(mediaPlayer IZegoMediaPlayer) {
	if mediaPlayer == nil {
		return
	}
	index := mediaPlayer.GetIndex()

	e.mediaPlayerLock.Lock()
	defer e.mediaPlayerLock.Unlock()
	if _, ok := e.mediaPlayerImplMap[index]; ok {
		C.zego_express_destroy_media_player(e.handle, C.enum_zego_media_player_instance_index(index))
		delete(e.mediaPlayerImplMap, index)
	}
}

func (e *engineImpl) CallExperimentalAPI(params string) string {
	cParams := StringToCString(params)
	defer FreeCString(cParams)
	var tempResult *C.char = nil
	C.zego_express_call_experimental_api(e.handle, cParams, &tempResult)
	result := C.GoString(tempResult)
	C.zego_express_free_call_experimental_api_result(e.handle, tempResult)
	return result
}

type mediaPlayerImpl struct {
	handle                C.zego_handle
	handlerLock           sync.Mutex
	eventHandler          IZegoMediaPlayerEventHandler
	audioHandler          IZegoMediaPlayerAudioHandler
	instanceIndex         int
	loadResourceCallbacks *list.List
	seekToCallbacks       map[int]ZegoMediaPlayerSeekToCallback
}

func (mediaPlayer *mediaPlayerImpl) SetEventHandler(handler IZegoMediaPlayerEventHandler) {
	mediaPlayer.handlerLock.Lock()
	defer mediaPlayer.handlerLock.Unlock()
	mediaPlayer.eventHandler = handler
}

func (mediaPlayer *mediaPlayerImpl) SetAudioHandler(handler IZegoMediaPlayerAudioHandler) {
	mediaPlayer.handlerLock.Lock()
	defer mediaPlayer.handlerLock.Unlock()
	mediaPlayer.audioHandler = handler
}

func (mediaPlayer *mediaPlayerImpl) LoadResource(path string, callback ZegoMediaPlayerLoadResourceCallback) {
	cPath := StringToCString(path)
	defer FreeCString(cPath)
	result := C.zego_express_media_player_load_resource(mediaPlayer.handle, cPath, C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex))
	if result != C.ZEGO_ERRCODE_COMMON_SUCCESS {
		if callback != nil {
			callback(int(result))
		}
		return
	}
	mediaPlayer.handlerLock.Lock()
	defer mediaPlayer.handlerLock.Unlock()
	mediaPlayer.loadResourceCallbacks.PushBack(callback)
}

func (mediaPlayer *mediaPlayerImpl) Start() {
	C.zego_express_media_player_start(mediaPlayer.handle, C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex))
}

func (mediaPlayer *mediaPlayerImpl) Stop() {
	C.zego_express_media_player_stop(mediaPlayer.handle, C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex))
}

func (mediaPlayer *mediaPlayerImpl) Pause() {
	C.zego_express_media_player_pause(mediaPlayer.handle, C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex))
}

func (mediaPlayer *mediaPlayerImpl) Resume() {
	C.zego_express_media_player_resume(mediaPlayer.handle, C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex))
}

func (mediaPlayer *mediaPlayerImpl) SeekTo(millisecond uint64, callback ZegoMediaPlayerSeekToCallback) {
	cSeq := C.zego_express_get_increase_seq(mediaPlayer.handle)
	mediaPlayer.handlerLock.Lock()
	mediaPlayer.seekToCallbacks[int(cSeq)] = callback
	mediaPlayer.handlerLock.Unlock()
	C.zego_express_media_player_seek_to(mediaPlayer.handle, C.ulonglong(millisecond), C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex), &cSeq)
}

func (mediaPlayer *mediaPlayerImpl) EnableRepeat(enable bool) {
	C.zego_express_media_player_enable_repeat(mediaPlayer.handle, C.bool(enable), C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex))
}

func (mediaPlayer *mediaPlayerImpl) EnableAux(enable bool) {
	C.zego_express_media_player_enable_aux(mediaPlayer.handle, C.bool(enable), C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex))
}

func (mediaPlayer *mediaPlayerImpl) SetVolume(volume int) {
	C.zego_express_media_player_set_volume(mediaPlayer.handle, C.int(volume), C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex))
}

func (mediaPlayer *mediaPlayerImpl) SetPlayVolume(volume int) {
	C.zego_express_media_player_set_play_volume(mediaPlayer.handle, C.int(volume), C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex))
}

func (mediaPlayer *mediaPlayerImpl) SetPublishVolume(volume int) {
	C.zego_express_media_player_set_publish_volume(mediaPlayer.handle, C.int(volume), C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex))
}

func (mediaPlayer *mediaPlayerImpl) GetPlayVolume() int {
	var volume C.int
	C.zego_express_media_player_get_play_volume(mediaPlayer.handle, C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex), &volume)
	return int(volume)
}

func (mediaPlayer *mediaPlayerImpl) GetPublishVolume() int {
	var volume C.int
	C.zego_express_media_player_get_publish_volume(mediaPlayer.handle, C.enum_zego_media_player_instance_index(mediaPlayer.instanceIndex), &volume)
	return int(volume)
}

func (mediaPlayer *mediaPlayerImpl) GetIndex() int {
	return mediaPlayer.instanceIndex
}

func createEngineInner(profile ZegoMultiEngineProfile, handler IZegoEventHandler) (*engineImpl, bool) {
	engine := NewEngineImpl()

	result := engine.init(profile, handler)
	if result != ZegoErrorCodeCommonSuccess {
		handler.OnDebugError(result, "CreateEngine", "CreateEngine failed")
		return engine, false
	}
	return engine, true
}

func createEngine(profile ZegoMultiEngineProfile, handler IZegoEventHandler) IZegoExpressEngine {
	result, ok := createEngineInner(profile, handler)

	if !ok {
		return nil
	}

	maxPublishChannelCount := 4
	if profile.AdvancedConfig != nil {
		if value, exists := profile.AdvancedConfig["max_publish_channels"]; exists {
			if count, err := strconv.Atoi(value); err == nil {
				maxPublishChannelCount = count
			}
		}
	}
	for i := 0; i < maxPublishChannelCount; i++ {
		C.zego_express_enable_camera(result.handle, C.bool(false), C.zego_exp_notify_device_state_mode_open, C.enum_zego_publish_channel(i))
	}
	return result
}

func destroyEngine(engine IZegoExpressEngine, callback ZegoDestroyCompletionCallback) {
	if engine == nil {
		return
	}
	if realEngine, ok := engine.(*engineImpl); ok {
		engineDestroyCallbackLock.Lock()
		// 考虑换成seq, 避免engineImpl的地址重复使用
		engineDestroyCallbacks[unsafe.Pointer(realEngine)] = callback
		engineDestroyCallbackLock.Unlock()
		C.zego_express_engine_uninit_async(realEngine.handle)
		realEngine.handlerLock.Lock()
		realEngine.eventHandler = nil
		realEngine.audioDataHandler = nil
		realEngine.customAudioProcessHandler = nil
		realEngine.handlerLock.Unlock()
		realEngine.pinner.Unpin()
	}
}

func setLogConfig(config ZegoLogConfig) {
	cLogConfig := C.struct_zego_log_config{
		log_size:  C.ulonglong(config.LogSize),
		log_count: C.uint(config.LogCount),
	}
	setCharArray(&cLogConfig.log_path[0], config.LogPath, C.ZEGO_EXPRESS_MAX_COMMON_LEN)
	C.zego_express_set_log_config(nil, cLogConfig)
}

func setCallbackEventHandler(handler IZegoCallbackEventHandler) {
	callbackEventLock.Lock()
	defer callbackEventLock.Unlock()
	callbackEventHandler = handler
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

type callbackHandler struct {
	callbackChan chan func()
}

func (h *callbackHandler) processLoop() {
	for callback := range h.callbackChan {
		callback()
	}
}

func (h *callbackHandler) dispatchInCallbackGoroutine(callbackFunc func()) {
	if h.callbackChan == nil {
		return
	}
	select {
	case h.callbackChan <- callbackFunc:
		return
	default:
		callbackEventLock.Lock()
		defer callbackEventLock.Unlock()
		if callbackEventHandler != nil {
			callbackEventHandler.OnCallbackDiscarded()
		}
	}
}

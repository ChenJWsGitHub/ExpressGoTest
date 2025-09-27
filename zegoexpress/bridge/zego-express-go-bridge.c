#include <stddef.h>
#include "zego-express-go-bridge.h"
#include "zego-express-engine.h"
#include "zego-express-room.h"
#include "zego-express-publisher.h"

// 声明由Go实现的函数
extern void GoOnApiCalledResult(int, const char*, const char*);
extern void GoLoginResultCallback(int, const char*, int);
extern void GoLogoutResultCallback(int, const char*, int);
extern void GoOnIMSendBroadcastMessageResult(zego_error, unsigned long long msg_id, int);
extern void GoOnPlayerAudioData(const unsigned char *, unsigned int, struct zego_audio_frame_param, const char *);

extern void GoOnDebugError(int error_code, const char* func_name, const char* info);
extern void GoOnRoomStateUpdate(const char *room_id, enum zego_room_state state, zego_error error_code, const char *extend_data);
extern void GoOnRoomStreamUpdate(const char *room_id, enum zego_update_type update_type, const struct zego_stream *stream_info_list, unsigned int stream_info_count, const char *extended_data);
extern void GoOnRoomStateChanged(const char *room_id, enum zego_room_state_changed_reason reason, zego_error error_code, const char *extended_data);
extern void GoOnRoomTokenWillExpire(const char *room_id, int remain_time_in_second);
extern void GoOnPublisherStateUpdate(const char *stream_id, enum zego_publisher_state state, zego_error error_code, const char *extend_data);
extern void GoOnPublisherQualityUpdate(const char *stream_id, struct zego_publish_stream_quality quality);
extern void GoOnPublisherStreamEvent(enum zego_stream_event event_id, const char *stream_id, const char *extra_info);
extern void GoOnPlayerStateUpdate(const char *stream_id, enum zego_player_state state, zego_error error_code, const char *extend_data);
extern void GoOnPlayerQualityUpdate(const char *stream_id, struct zego_play_stream_quality quality);
extern void GoOnPlayerRecvSei(struct zego_media_side_info info);
extern void GoOnPlayerStreamEvent(enum zego_stream_event event_id, const char *stream_id, const char *extra_info);

static void bridge_go_on_api_called_result(int error_code, const char *func_name, const char *info, void *user_context) {
    GoOnApiCalledResult(error_code, func_name, info);
}

static void bridge_go_login_callback(zego_error code, const char *ext_data, const char *room_id, zego_seq seq, void *ctx) {
    GoLoginResultCallback(code, ext_data, seq);
}

static void bridge_go_logout_callback(zego_error code, const char *ext_data, const char *room_id, zego_seq seq, void *ctx) {
    GoLogoutResultCallback(code, ext_data, seq);
}

static void bridge_go_on_im_send_broadcast_message_result(const char *room_id, unsigned long long message_id, zego_error error_code, zego_seq seq, void *user_context) {
    GoOnIMSendBroadcastMessageResult(error_code, message_id, seq);
}

static void bridge_go_on_player_audio_data(const unsigned char *data, unsigned int data_length, struct zego_audio_frame_param param, const char *stream_id, void *user_context) {
    GoOnPlayerAudioData(data, data_length, param, stream_id);
}

static void bridge_go_on_debug_error(int error_code, const char *func_name, const char *info, void *user_context) {
    GoOnDebugError(error_code, func_name, info);
}

static void bridge_go_on_room_state_update(const char *room_id, enum zego_room_state state, zego_error error_code, const char *extend_data, void *user_context) {
    GoOnRoomStateUpdate(room_id, state, error_code, extend_data);
}

static void bridge_go_on_room_stream_update(const char *room_id, enum zego_update_type update_type, const struct zego_stream *stream_info_list, unsigned int stream_info_count, const char *extended_data, void *user_context) {
    GoOnRoomStreamUpdate(room_id, update_type, stream_info_list, stream_info_count, extended_data);
}

static void bridge_go_on_room_state_changed(const char *room_id, enum zego_room_state_changed_reason reason, zego_error error_code, const char *extended_data, void *user_context) {
    GoOnRoomStateChanged(room_id, reason, error_code, extended_data);
}

static void bridge_go_on_room_token_will_expire(const char *room_id, int remain_time_in_second, void *user_context) {
    GoOnRoomTokenWillExpire(room_id, remain_time_in_second);
}

static void bridge_go_on_publisher_state_update(const char *stream_id, enum zego_publisher_state state, zego_error error_code, const char *extend_data, void *user_context) {
    GoOnPublisherStateUpdate(stream_id, state, error_code, extend_data);
}

static void bridge_go_on_publisher_quality_update(const char *stream_id, struct zego_publish_stream_quality quality, void *user_context) {
    GoOnPublisherQualityUpdate(stream_id, quality);
}

static void bridge_go_on_publisher_stream_event(enum zego_stream_event event_id, const char *stream_id, const char *extra_info, void *user_context) {
    GoOnPublisherStreamEvent(event_id, stream_id, extra_info);
}

static void bridge_go_on_player_state_update(const char *stream_id, enum zego_player_state state, zego_error error_code, const char *extend_data, void *user_context) {
    GoOnPlayerStateUpdate(stream_id, state, error_code, extend_data);
}

static void bridge_go_on_player_quality_update(const char *stream_id, struct zego_play_stream_quality quality, void *user_context) {
    GoOnPlayerQualityUpdate(stream_id, quality);
}

static void bridge_go_on_player_recv_sei(struct zego_media_side_info info, void *user_context) {
    GoOnPlayerRecvSei(info);
}

static void bridge_go_on_player_stream_event(enum zego_stream_event event_id, const char *stream_id, const char *extra_info, void *user_context) {
    GoOnPlayerStreamEvent(event_id, stream_id, extra_info);
}

void zego_express_go_bridge_init() {
    zego_register_api_called_result_callback(bridge_go_on_api_called_result, NULL);
    zego_register_room_login_result_callback(bridge_go_login_callback, NULL);
    zego_register_room_logout_result_callback(bridge_go_logout_callback, NULL);
    zego_register_im_send_broadcast_message_result_callback(bridge_go_on_im_send_broadcast_message_result, NULL);
    zego_register_player_audio_data_callback(bridge_go_on_player_audio_data, NULL);
    zego_register_debug_error_callback(bridge_go_on_debug_error, NULL);
    zego_register_room_state_update_callback(bridge_go_on_room_state_update, NULL);
    zego_register_room_stream_update_callback(bridge_go_on_room_stream_update, NULL);
    zego_register_room_state_changed_callback(bridge_go_on_room_state_changed, NULL);
    zego_register_room_token_will_expire_callback(bridge_go_on_room_token_will_expire, NULL);
    zego_register_publisher_state_update_callback(bridge_go_on_publisher_state_update, NULL);
    zego_register_publisher_quality_update_callback(bridge_go_on_publisher_quality_update, NULL);
    zego_register_publisher_stream_event_callback(bridge_go_on_publisher_stream_event, NULL);
    zego_register_player_state_update_callback(bridge_go_on_player_state_update, NULL);
    zego_register_player_quality_update_callback(bridge_go_on_player_quality_update, NULL);
    zego_register_player_recv_media_side_info_callback(bridge_go_on_player_recv_sei, NULL);
    zego_register_player_stream_event_callback(bridge_go_on_player_stream_event, NULL);
}

#ifndef __ZEGO_EXPRESS_CUSTOM_AUDIO_SOURCE_H__
#define __ZEGO_EXPRESS_CUSTOM_AUDIO_SOURCE_H__

#include "zego-express-defines.h"

ZEGO_BEGIN_DECLS
/// Create custom audio source instance.
///
/// Available since: 3.25.0
/// Description: Create custom audio source instance.
/// Use cases: Typically used to create an independent custom audio source and mix it into the publish stream as a mix source.
/// When to call: After the engine is created [createEngine].
/// Restrictions: None.
/// Caution:
///  1. ZegoCustomAudioSource can only be used as a mix source, not as the main capture source.
///  2. Currently, up to 2 custom audio source instances are supported, with the screen sharing audio source occupying one instance.
///
/// @return Custom audio source instance.
#ifndef ZEGOEXP_EXPLICIT
ZEGOEXP_API int EXP_CALL zego_express_create_custom_audio_source(zego_handle handle);
#else
typedef int(EXP_CALL *pfnzego_express_create_custom_audio_source)(zego_handle handle);
#endif

/// Destroy custom audio source instance.
///
/// Available since: 3.25.0
/// Description: Destroy custom audio source instance.
/// Use cases: Typically used to destroy a custom audio mix source instance.
/// When to call: After the engine is created [createEngine].
/// Restrictions: None.
/// Caution: ZegoCustomAudioSource can only be used as a mix source, not as the main capture source.
///
/// @param audio_source_index Custom audio source instance.
#ifndef ZEGOEXP_EXPLICIT
ZEGOEXP_API zego_error EXP_CALL zego_express_destroy_custom_audio_source(zego_handle handle,
                                                                         int audio_source_index);
#else
typedef zego_error(EXP_CALL *pfnzego_express_destroy_custom_audio_source)(zego_handle handle,
                                                                          int audio_source_index);
#endif

/// Start custom audio source capture.
///
/// Available since: 3.25.0
/// Description: Start custom audio source capture.
/// Use cases: Typically used in educational scenarios that require switching between different audio capture sources.
/// When to call: After the custom audio source is created by [createCustomAudioSource].
/// Restrictions: None.
/// Caution: ZegoCustomAudioSource can only be used as a mix source, not as the main capture source.
///
/// @return The result of the API call.
#ifndef ZEGOEXP_EXPLICIT
ZEGOEXP_API zego_error EXP_CALL
zego_express_custom_audio_source_start_capture(zego_handle handle, int audio_source_index);
#else
typedef zego_error(EXP_CALL *pfnzego_express_custom_audio_source_start_capture)(
    zego_handle handle, int audio_source_index);
#endif

/// Stop custom audio source capture.
///
/// Available since: 3.25.0
/// Description: Stop custom audio source capture.
/// Use cases: Typically used in educational scenarios that require switching between different audio capture sources.
/// When to call: After the custom audio source is created by [createCustomAudioSource].
/// Restrictions: None.
/// Caution: ZegoCustomAudioSource can only be used as a mix source, not as the main capture source.
///
/// @return The result of the API call.
#ifndef ZEGOEXP_EXPLICIT
ZEGOEXP_API zego_error EXP_CALL
zego_express_custom_audio_source_stop_capture(zego_handle handle, int audio_source_index);
#else
typedef zego_error(EXP_CALL *pfnzego_express_custom_audio_source_stop_capture)(
    zego_handle handle, int audio_source_index);
#endif

/// Enable custom audio source echo cancellation.
///
/// Available since: 3.25.0
/// Description: Enable custom audio source echo cancellation.
/// Use cases: Typically used in educational scenarios that require switching between different audio capture sources.
/// When to call: After the custom audio source is created by [createCustomAudioSource].
/// Restrictions: None.
/// Caution: ZegoCustomAudioSource can only be used as a mix source, not as the main capture source.
///
/// @param enable Whether to enable echo cancellation.
/// @return The result of the API call.
#ifndef ZEGOEXP_EXPLICIT
ZEGOEXP_API zego_error EXP_CALL zego_express_custom_audio_source_enable_aec(zego_handle handle,
                                                                            int audio_source_index,
                                                                            bool enable);
#else
typedef zego_error(EXP_CALL *pfnzego_express_custom_audio_source_enable_aec)(zego_handle handle,
                                                                             int audio_source_index,
                                                                             bool enable);
#endif

/// Set custom audio source volume.
///
/// Available since: 3.25.0
/// Description: Set the volume of the custom audio source.
/// Use cases: Typically used in scenarios that require switching between different audio capture sources.
/// When to call: After the custom audio source is created by [createCustomAudioSource].
/// Restrictions: The volume range is 0 ~ 100.
/// Caution: ZegoCustomAudioSource can only be used as a mix source, not as the main capture source.
///
/// @param volume Volume, the range is 0 ~ 100.
/// @return The result of the API call.
#ifndef ZEGOEXP_EXPLICIT
ZEGOEXP_API zego_error EXP_CALL zego_express_custom_audio_source_set_volume(zego_handle handle,
                                                                            int audio_source_index,
                                                                            int volume);
#else
typedef zego_error(EXP_CALL *pfnzego_express_custom_audio_source_set_volume)(zego_handle handle,
                                                                             int audio_source_index,
                                                                             int volume);
#endif

/// Mute or unmute the custom audio source.
///
/// Available since: 3.25.0
/// Description: Mute or unmute the audio data of the custom audio source.
/// Use cases: Typically used in scenarios that require switching between different audio capture sources.
/// When to call: After the custom audio source is created by [createCustomAudioSource].
/// Restrictions: None.
/// Caution: ZegoCustomAudioSource can only be used as a mix source, not as the main capture source.
///
/// @param mute Whether to mute the custom audio source.
/// @return The result of the API call.
#ifndef ZEGOEXP_EXPLICIT
ZEGOEXP_API zego_error EXP_CALL zego_express_custom_audio_source_mute(zego_handle handle,
                                                                      int audio_source_index,
                                                                      bool mute);
#else
typedef zego_error(EXP_CALL *pfnzego_express_custom_audio_source_mute)(zego_handle handle,
                                                                       int audio_source_index,
                                                                       bool mute);
#endif

/// Push audio frame to the custom audio source.
///
/// Available since: 3.25.0
/// Description: Push audio frame to the custom audio source.
/// Use cases: Typically used when developers manage audio data capture themselves and input audio frames to a specified custom audio source.
/// When to call: After the custom audio source is created by [createCustomAudioSource].
/// Restrictions: None.
/// Caution: ZegoCustomAudioSource can only be used as a mix source, not as the main capture source.
///
/// @param data PCM buffer data.
/// @param data_length The total length of the buffer data.
/// @param param The param of this PCM audio frame.
/// @return The result of the API call.
#ifndef ZEGOEXP_EXPLICIT
ZEGOEXP_API zego_error EXP_CALL zego_express_custom_audio_source_push_audio_frame(
    zego_handle handle, int audio_source_index, unsigned char *data, unsigned int data_length,
    struct zego_audio_frame_param param);
#else
typedef zego_error(EXP_CALL *pfnzego_express_custom_audio_source_push_audio_frame)(
    zego_handle handle, int audio_source_index, unsigned char *data, unsigned int data_length,
    struct zego_audio_frame_param param);
#endif

ZEGO_END_DECLS

#endif

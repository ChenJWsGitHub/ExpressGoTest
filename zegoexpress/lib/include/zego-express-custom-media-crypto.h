
#ifndef __ZEGO_EXPRESS_CUSTOM_MEDIA_CRYPTO_H__
#define __ZEGO_EXPRESS_CUSTOM_MEDIA_CRYPTO_H__

#include "zego-express-defines.h"

ZEGO_BEGIN_DECLS
/// Enable custom media crypto feature.
///
/// Available since: 3.24.0
/// Description: Enable custom media crypto feature.
/// When to call: After [createEngine] .
/// Caution: Once this feature is enabled for an engine instance, it is not recommended to disable it. Disabling requires stopping the engine before the operation can be invoked.
///
/// @param enable enable or disable
#ifndef ZEGOEXP_EXPLICIT
ZEGOEXP_API zego_error EXP_CALL zego_express_enable_custom_media_crypto(zego_handle handle,
                                                                        bool enable);
#else
typedef zego_error(EXP_CALL *pfnzego_express_enable_custom_media_crypto)(zego_handle handle,
                                                                         bool enable);
#endif

/// Get the length of the encrypted target.
///
/// Available since: 3.24.0
/// Description: Get the length of the encrypted destination.
/// When to Trigger: Encrypted per package.
///
/// @param s_rc_len Length of the source.
/// @param instance The context returned by the callback [createMediaCrypto].
/// @param user_context Context of user.
/// @return Length of the destination.
typedef int (*zego_on_custom_media_crypto_get_encrypt_dst_len)(zego_handle handle, int s_rc_len,
                                                               void *instance, void *user_context);

#ifndef ZEGOEXP_EXPLICIT
ZEGOEXP_API void EXP_CALL zego_register_custom_media_crypto_get_encrypt_dst_len_callback(
    zego_handle handle, zego_on_custom_media_crypto_get_encrypt_dst_len callback_func,
    void *user_context);
#else
typedef int(EXP_CALL *pfnzego_register_custom_media_crypto_get_encrypt_dst_len_callback)(
    zego_handle handle, zego_on_custom_media_crypto_get_encrypt_dst_len callback_func,
    void *user_context);
#endif

/// Audio encryption callback.
///
/// Available since: 3.24.0
/// Description: Audio encryption callback.
/// When to Trigger: Encrypt each audio package.
///
/// @param in_buffer The block data of the input data.
/// @param in_buffer_size Length of inBuffer.
/// @param out_buffer The block data of the output data.
/// @param out_buffer_size Length of outBuffer.
/// @param instance The context returned by the callback [createMediaCrypto].
/// @param user_context Context of user.
/// @return Returns 0 on success, non-zero on failure.
typedef int (*zego_on_custom_media_crypto_audio_encrypt)(
    zego_handle handle, const unsigned char *in_buffer, int in_buffer_size,
    unsigned char *out_buffer, int out_buffer_size, void *instance, void *user_context);

#ifndef ZEGOEXP_EXPLICIT
ZEGOEXP_API void EXP_CALL zego_register_custom_media_crypto_audio_encrypt_callback(
    zego_handle handle, zego_on_custom_media_crypto_audio_encrypt callback_func,
    void *user_context);
#else
typedef int(EXP_CALL *pfnzego_register_custom_media_crypto_audio_encrypt_callback)(
    zego_handle handle, zego_on_custom_media_crypto_audio_encrypt callback_func,
    void *user_context);
#endif

/// Video encryption callback.
///
/// Available since: 3.24.0
/// Description: Video encryption callback.
/// When to Trigger: Encrypt each video package.
///
/// @param in_buffer The block data of the input data.
/// @param in_buffer_size Length of inBuffer.
/// @param out_buffer The block data of the output data.
/// @param out_buffer_size Length of outBuffer.
/// @param instance The context returned by the callback [createMediaCrypto].
/// @param user_context Context of user.
/// @return Returns 0 on success, non-zero on failure.
typedef int (*zego_on_custom_media_crypto_video_encrypt)(
    zego_handle handle, const unsigned char *in_buffer, int in_buffer_size,
    unsigned char *out_buffer, int out_buffer_size, void *instance, void *user_context);

#ifndef ZEGOEXP_EXPLICIT
ZEGOEXP_API void EXP_CALL zego_register_custom_media_crypto_video_encrypt_callback(
    zego_handle handle, zego_on_custom_media_crypto_video_encrypt callback_func,
    void *user_context);
#else
typedef int(EXP_CALL *pfnzego_register_custom_media_crypto_video_encrypt_callback)(
    zego_handle handle, zego_on_custom_media_crypto_video_encrypt callback_func,
    void *user_context);
#endif

/// Audio decryption callback.
///
/// Available since: 3.24.0
/// Description: Audio decryption callback.
/// When to Trigger: Decrypt each audio package.
///
/// @param in_buffer The block data of the input data.
/// @param in_buffer_size Length of inBuffer.
/// @param out_buffer The block data of the output data.
/// @param out_buffer_size Length of outBuffer.
/// @param instance The context returned by the callback [createMediaCrypto].
/// @param user_context Context of user.
/// @return Returns 0 on success, non-zero on failure.
typedef int (*zego_on_custom_media_crypto_audio_decrypt)(
    zego_handle handle, const unsigned char *in_buffer, int in_buffer_size,
    unsigned char *out_buffer, int out_buffer_size, void *instance, void *user_context);

#ifndef ZEGOEXP_EXPLICIT
ZEGOEXP_API void EXP_CALL zego_register_custom_media_crypto_audio_decrypt_callback(
    zego_handle handle, zego_on_custom_media_crypto_audio_decrypt callback_func,
    void *user_context);
#else
typedef int(EXP_CALL *pfnzego_register_custom_media_crypto_audio_decrypt_callback)(
    zego_handle handle, zego_on_custom_media_crypto_audio_decrypt callback_func,
    void *user_context);
#endif

/// Video decryption callback.
///
/// Available since: 3.24.0
/// Description: Video decryption callback.
/// When to Trigger: Decrypt each video package.
///
/// @param in_buffer The block data of the input data.
/// @param in_buffer_size Length of inBuffer.
/// @param out_buffer The block data of the output data.
/// @param out_buffer_size Length of outBuffer.
/// @param instance The context returned by the callback [createMediaCrypto].
/// @param user_context Context of user.
/// @return Returns 0 on success, non-zero on failure.
typedef int (*zego_on_custom_media_crypto_video_decrypt)(
    zego_handle handle, const unsigned char *in_buffer, int in_buffer_size,
    unsigned char *out_buffer, int out_buffer_size, void *instance, void *user_context);

#ifndef ZEGOEXP_EXPLICIT
ZEGOEXP_API void EXP_CALL zego_register_custom_media_crypto_video_decrypt_callback(
    zego_handle handle, zego_on_custom_media_crypto_video_decrypt callback_func,
    void *user_context);
#else
typedef int(EXP_CALL *pfnzego_register_custom_media_crypto_video_decrypt_callback)(
    zego_handle handle, zego_on_custom_media_crypto_video_decrypt callback_func,
    void *user_context);
#endif

/// Create a custom media crypto object.
///
/// Available since: 3.24.0
/// Description: The SDK notifies that media data encryption/decryption needs to begin.
/// When to Trigger: After calling [startPlayingStream] or [startPublishingStream] successfully.
///
/// @param stream_id Stream ID.
/// @param is_encrypt Whether to encrypt. true indicates encryption, false indicates decryption.
/// @param user_context Context of user.
/// @return Custom media crypto object.
typedef void *(*zego_on_custom_media_crypto_create)(zego_handle handle, const char *stream_id,
                                                    bool is_encrypt, void *user_context);

#ifndef ZEGOEXP_EXPLICIT
ZEGOEXP_API void EXP_CALL zego_register_custom_media_crypto_create_callback(
    zego_handle handle, zego_on_custom_media_crypto_create callback_func, void *user_context);
#else
typedef void *(EXP_CALL *pfnzego_register_custom_media_crypto_create_callback)(
    zego_handle handle, zego_on_custom_media_crypto_create callback_func, void *user_context);
#endif

/// Destroy a custom media crypto object.
///
/// Available since: 3.24.0
/// Description: Destroy a custom media crypto object.
/// When to Trigger: At the end of a stream.
///
/// @param instance The custom media crypto object returned by the callback [createMediaCrypto].
/// @param user_context Context of user.
typedef void (*zego_on_custom_media_crypto_destroy)(zego_handle handle, void *instance,
                                                    void *user_context);

#ifndef ZEGOEXP_EXPLICIT
ZEGOEXP_API void EXP_CALL zego_register_custom_media_crypto_destroy_callback(
    zego_handle handle, zego_on_custom_media_crypto_destroy callback_func, void *user_context);
#else
typedef void(EXP_CALL *pfnzego_register_custom_media_crypto_destroy_callback)(
    zego_handle handle, zego_on_custom_media_crypto_destroy callback_func, void *user_context);
#endif

ZEGO_END_DECLS

#endif

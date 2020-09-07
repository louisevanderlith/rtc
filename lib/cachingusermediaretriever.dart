library cachingusermediaretriever;

import "dart:async";
import "dart:html";

class CachingUserMediaRetriever {
  MediaStream mediaStream;

  Future<MediaStream> get() {
    Completer completer = new Completer<MediaStream>();
    if(mediaStream != null)
      completer.complete(mediaStream);
    else
      window.navigator.getUserMedia(audio: false, video: true).then((MediaStream stream) {
        mediaStream = stream;
        completer.complete(mediaStream);
      });
    return completer.future;
  }
}
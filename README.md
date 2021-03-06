# VAD en GO
Copyright 2015 Thomas Solignac

# Purpose
Voice Activity Detection aims to detect the start/end of speeches (getting complete sentences) from an audio stream. 

# Tools

Install before :
apt-get install libsndfile1-dev

Lib SND :
https://github.com/mkb218/gosndfile

# Samples

Test in data are generated by Audacity

.wav : Not working "Unknown format"   
.aiff : OK   

# Knowledge

```cpp
type Info struct {
	Frames     int64
	// An audio frame, or sample, contains amplitude (loudness) information at that particular point in time.
	// To produce sound, tens of thousands of frames are played in sequence to produce frequencies.
	
	Samplerate int32
	// Number of samples of audio carried per second, measured in Hz or kHz (one kHz being 1 000 Hz).
	// For example, 44 100 samples per second can be expressed as either 44 100 Hz, or 44.1 kHz.
	
	Channels   int32
	// An audio channel or audio track is a audio signal communications channel in a storage device,
	// used in operations such as multi-track recording and sound reinforcement.
	// In stereo they are twho channels : Left and Right
	
	Format     Format
	Sections   int32
	
	Seekable   int32
	// Buffered zone that you can pick ?
}
```


About audio encoding itself :
http://www.commentcamarche.net/contents/81-le-son-numerique

And the details :
http://en.wikipedia.org/wiki/Pulse-code_modulation

The channels are encoded : LR LR LR LR
http://csserver.evansville.edu/~blandfor/EE356/WavFormatDocs.pdf

Example os use :
http://blog.afandian.com/2012/07/sound-file-plotter-in-go-using-gosndfile-libsndfile/

# Algorithms

"voice segmentation algorithm"
"voice activity detection"

Audio Engineering Society
http://www.aes.org/

## Exact problem

Entropy Method :
https://www.google.fr/url?sa=t&rct=j&q=&esrc=s&source=web&cd=2&cad=rja&uact=8&ved=0CC8QFjAB&url=http%3A%2F%2Fresearchgroups.msu.edu%2Fsystem%2Ffiles%2Fgroup-publication%2Fmwscas02_entropy_endpoint_final.pdf&ei=Xh0GVfHELsv1UMiYhMgF&usg=AFQjCNFrI72ssEOnGZeg2jT-pQ-LBK4nsw&sig2=4c2AoDF4QmNxpUySIIwzUw

Fourrier Method :
http://www.clsp.jhu.edu/~zak/is-icassp03.pdf
http://www.ijcta.com/documents/volumes/vol4issue2/ijcta2013040202.pdf

K-nearest-neighbor (KNN) and linear spectral pairs-vector quantization (LSP-VQ)
http://www.cs.bc.edu/~hjiang/papers/journal/tsap.pdf

Schur adaptive filtering
https://www.google.fr/url?sa=t&rct=j&q=&esrc=s&source=web&cd=13&cad=rja&uact=8&ved=0CJ8BEBYwDA&url=https%3A%2F%2Fwww.amcs.uz.zgora.pl%2F%3Faction%3Ddownload%26pdf%3DAMCS_2014_24_2_3.pdf&ei=UWsHVan2C9PjaqSKgagE&usg=AFQjCNGa4eJ1ZR20TDMRtAb2KJfjliBHBQ&sig2=_IMv9pBg1yKMgSqtJIGJNg

Tuto :
http://practicalcryptography.com/miscellaneous/machine-learning/voice-activity-detection-vad-tutorial/

## Similar

Word/Sub-word/Syllable segmentation
http://www.asel.udel.edu/icslp/cdrom/vol2/439/a439.pdf
http://esatjournals.org/Volumes/IJRET/2013V02/I09/IJRET20130209061.pdf
http://lantana.tenet.res.in/website_files/publications/Speech/speech_communication_o.pdf

## Helps

Maybe to use later ? For prediction
http://en.wikipedia.org/wiki/Forward%E2%80%93backward_algorithm

PPT about noise (exactly my problem)
https://www.google.fr/url?sa=t&rct=j&q=&esrc=s&source=web&cd=1&cad=rja&uact=8&ved=0CCcQFjAA&url=http%3A%2F%2Fwww.ece.lsu.edu%2Fwu%2Fpublic_html%2Fdemo_end_of_speech%2Fdemo_end_of_speech.ppt&ei=lBwGVaTNGoStUZj6gdgH&usg=AFQjCNEIcwr2qgZEzJ4rGKLZe_Ouk4Cmzw&sig2=RnfjXljK9N5-pdLKBZ3A3g

Vocabulary
http://en.wikipedia.org/wiki/Talkspurt


## Already made solution

http://libvad.com/

## With real time

WebRTC does that ?
https://code.google.com/p/webrtc/source/browse/trunk/#trunk%2Fwebrtc%2Fcommon_audio%2Fvad

Paper "simple"
http://ms12.voip.edu.tw/~paul/Papper/Steganography/iLBC/(VAD)Real-Time_VAD_Algorithm.pdf


# API Speech Recognition

Google Speech API
https://github.com/gillesdemey/google-speech-v2

MIT Hosting SR API
https://code.google.com/p/wami/

HP Solution
https://www.idolondemand.com/developer/apis/recognizespeech#overview


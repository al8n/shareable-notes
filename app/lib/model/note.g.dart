// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'note.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

Note _$NoteFromJson(Map<String, dynamic> json) {
  return Note(
    oid: json['_id'] as String,
    id: json['id'] as String,
    title: json['title'] as String,
    createTime: json['createtime'] as String,
    modifyTime: json['modifytime'] as String,
    highlightColor:
        const ColorConverter().fromJson(json['highlight_style'] as String),
    topics: (json['topic'] as List<dynamic>)
        .map(
            (e) => e == null ? null : Topic.fromJson(e as Map<String, dynamic>))
        .toList(),
    mindlinks:
        (json['mindlinks'] as List<dynamic>).map((e) => e as String?).toList(),
    meta: MetaData.fromJson(json['metadata'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$NoteToJson(Note instance) => <String, dynamic>{
      '_id': instance.oid,
      'id': instance.id,
      'mindlinks': instance.mindlinks,
      'createtime': instance.createTime,
      'modifytime': instance.modifyTime,
      'highlight_style': const ColorConverter().toJson(instance.highlightColor),
      'title': instance.title,
      'topic': instance.topics,
      'metadata': instance.meta,
    };

Topic _$TopicFromJson(Map<String, dynamic> json) {
  return Topic(
    text: json['text'] as String?,
    audio: json['audio'] as String?,
    photo: json['photo'] as String?,
    tag: json['tag'] as String?,
    htext: json['htext'] as String?,
    haudio: json['haudio'] as String?,
    hpic: json['hpic'] == null
        ? null
        : HPicture.fromJson(json['hpic'] as Map<String, dynamic>),
    hvideo: json['hvideo'] as String?,
  );
}

Map<String, dynamic> _$TopicToJson(Topic instance) => <String, dynamic>{
      'text': instance.text,
      'tag': instance.tag,
      'audio': instance.audio,
      'photo': instance.photo,
      'htext': instance.htext,
      'haudio': instance.haudio,
      'hvideo': instance.hvideo,
      'hpic': instance.hpic,
    };

HPicture _$HPictureFromJson(Map<String, dynamic> json) {
  return HPicture(
    picture: json['pic'] as String?,
    text: json['text'] as String?,
  );
}

Map<String, dynamic> _$HPictureToJson(HPicture instance) => <String, dynamic>{
      'pic': instance.picture,
      'text': instance.text,
    };

MetaData _$MetaDataFromJson(Map<String, dynamic> json) {
  return MetaData(
    id: json['id'] as String,
    title: json['title'] as String,
  );
}

Map<String, dynamic> _$MetaDataToJson(MetaData instance) => <String, dynamic>{
      'id': instance.id,
      'title': instance.title,
    };

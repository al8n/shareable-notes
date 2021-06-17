import 'dart:ui';

import 'package:json_annotation/json_annotation.dart';

part 'note.g.dart';

extension RGB on String {
  Color toColor() {
    if (this.length == 0) {
      return Color(0xffffffff);
    }
    List<String> values = this.trim().replaceAll("rgb", "").replaceAll("(", "").replaceAll(")", "").split(",");
    return Color.fromRGBO(int.parse(values[0]), int.parse(values[1]), int.parse(values[2]), 1.0);
  }
}

class ColorConverter implements JsonConverter<Color, String> {
  const ColorConverter();

  @override
  Color fromJson(String json) {
    return json.toColor();
  }

  @override
  String toJson(Color json) => "rgb(${json.red},${json.green},${json.blue})";
}

@JsonSerializable()
class Note {
  @JsonKey(name: "_id")
  final String oid;

  @JsonKey(name: "rootNote", defaultValue: <Node>[])
  final List<Node?> rootNodes;

  @JsonKey(name: "summaryNote", defaultValue: <Map<String, List<String>>>{})
  final Map<String, List<String>> summaryNodes;

  @JsonKey(name: "linkNote", defaultValue: <Map<String, String>>{})
  final Map<String, List<String>> linkNodes;

  @JsonKey(name: "metadata")
  final MetaData meta;

  Note({
    required this.oid,
    required this.meta,
    required this.rootNodes,
    required this.summaryNodes,
    required this.linkNodes,
  });

  factory Note.fromJson(Map<String, dynamic> json) => _$NoteFromJson(json);

  Map<String, dynamic> toJson() => _$NoteToJson(this);
}

@JsonSerializable()
@ColorConverter()
class Node {
  final String id;

  @JsonKey(name: "createtime")
  final String createTime;

  @JsonKey(name: "modifytime")
  final String modifyTime;

  @JsonKey(name: "highlight_style", defaultValue: "")
  final Color highlightColor;

  final String title;

  @JsonKey(name: "topic", defaultValue: <Topic>[])
  final List<Topic?> topics;

  @JsonKey(name: "mindlinks", defaultValue: <Node>[])
  final List<Node?> children;

  Node({
    required this.id,
    required this.createTime,
    required this.modifyTime,
    required this.title,
    required this.highlightColor,
    required this.topics,
    required this.children,
  });
}

@JsonSerializable()
class Topic {
  final String? text;
  final String? tag;
  final String? audio;
  final String? photo;
  final String? htext;
  final String? haudio;
  final String? hvideo;
  final HPicture? hpic;

  Topic({
    this.text,
    this.audio,
    this.photo,
    this.tag,
    this.htext,
    this.haudio,
    this.hpic,
    this.hvideo,
  });

  factory Topic.fromJson(Map<String, dynamic> json) => _$TopicFromJson(json);

  Map<String, dynamic> toJson() => _$TopicToJson(this);
}

@JsonSerializable()
class HPicture {
  @JsonKey(name: "pic")
  final String? picture;
  final String? text;

  HPicture({
    this.picture,
    this.text,
  });

  factory HPicture.fromJson(Map<String, dynamic> json) => _$HPictureFromJson(json);

  Map<String, dynamic> toJson() => _$HPictureToJson(this);
}


@JsonSerializable()
class MetaData {
  final String id;
  final String title;

  MetaData({
    required this.id,
    required this.title,
  });

  factory MetaData.fromJson(Map<String, dynamic> json) => _$MetaDataFromJson(json);

  Map<String, dynamic> toJson() => _$MetaDataToJson(this);
}
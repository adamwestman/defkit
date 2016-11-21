components {
  id: "on_create"
  component: "/defkit/events/on_create.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
components {
  id: "play_sound"
  component: "/defkit/actions/main1/play_sound.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  properties {
    id: "event"
    value: "#on_create"
    type: PROPERTY_TYPE_URL
  }
  properties {
    id: "delay_val"
    value: "1.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "gain_val"
    value: "0.5"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "gain_var"
    value: "score"
    type: PROPERTY_TYPE_HASH
  }
}
embedded_components {
  id: "sound"
  type: "sound"
  data: "sound: \"/main/sounds/strings2.wav\"\n"
  "looping: 1\n"
  "group: \"master\"\n"
  "gain: 1.0\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}

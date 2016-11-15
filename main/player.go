components {
  id: "destroy_point"
  component: "/defkit/actions/basic/destroy.script"
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
    value: "#on_point_collision"
    type: PROPERTY_TYPE_URL
  }
  properties {
    id: "target_type"
    value: "point"
    type: PROPERTY_TYPE_HASH
  }
}
components {
  id: "draw_score"
  component: "/defkit/actions/basic2/draw_score.script"
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
    value: "#on_step"
    type: PROPERTY_TYPE_URL
  }
}
components {
  id: "move_down"
  component: "/defkit/actions/move/move_fixed.script"
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
    value: "#on_down_pressed"
    type: PROPERTY_TYPE_URL
  }
  properties {
    id: "S"
    value: "true"
    type: PROPERTY_TYPE_BOOLEAN
  }
  properties {
    id: "speed_val"
    value: "100.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "target_self"
    value: "true"
    type: PROPERTY_TYPE_BOOLEAN
  }
}
components {
  id: "move_left"
  component: "/defkit/actions/move/move_fixed.script"
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
    value: "#on_left_pressed"
    type: PROPERTY_TYPE_URL
  }
  properties {
    id: "W"
    value: "true"
    type: PROPERTY_TYPE_BOOLEAN
  }
  properties {
    id: "speed_val"
    value: "100.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "target_self"
    value: "true"
    type: PROPERTY_TYPE_BOOLEAN
  }
}
components {
  id: "move_right"
  component: "/defkit/actions/move/move_fixed.script"
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
    value: "#on_right_pressed"
    type: PROPERTY_TYPE_URL
  }
  properties {
    id: "E"
    value: "true"
    type: PROPERTY_TYPE_BOOLEAN
  }
  properties {
    id: "speed_val"
    value: "100.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "target_self"
    value: "true"
    type: PROPERTY_TYPE_BOOLEAN
  }
}
components {
  id: "move_up"
  component: "/defkit/actions/move/move_fixed.script"
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
    value: "#on_up_pressed"
    type: PROPERTY_TYPE_URL
  }
  properties {
    id: "N"
    value: "true"
    type: PROPERTY_TYPE_BOOLEAN
  }
  properties {
    id: "speed_val"
    value: "100.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "target_self"
    value: "true"
    type: PROPERTY_TYPE_BOOLEAN
  }
}
components {
  id: "obj1"
  component: "/defkit/object.script"
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
    id: "type"
    value: "player"
    type: PROPERTY_TYPE_HASH
  }
}
components {
  id: "on_down_pressed"
  component: "/defkit/events/input/on_button_press.script"
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
    id: "action_id"
    value: "down"
    type: PROPERTY_TYPE_HASH
  }
}
components {
  id: "on_left_pressed"
  component: "/defkit/events/input/on_button_press.script"
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
    id: "action_id"
    value: "left"
    type: PROPERTY_TYPE_HASH
  }
}
components {
  id: "on_point_collision"
  component: "/defkit/events/contact/on_collision.script"
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
    id: "collision_group"
    value: "point"
    type: PROPERTY_TYPE_HASH
  }
}
components {
  id: "on_right_pressed"
  component: "/defkit/events/input/on_button_press.script"
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
    id: "action_id"
    value: "right"
    type: PROPERTY_TYPE_HASH
  }
}
components {
  id: "on_step"
  component: "/defkit/events/time/on_step.script"
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
  id: "on_up_pressed"
  component: "/defkit/events/input/on_button_press.script"
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
    id: "action_id"
    value: "up"
    type: PROPERTY_TYPE_HASH
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"\"\n"
  "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"player\"\n"
  "mask: \"default\"\n"
  "mask: \"point\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "    }\n"
  "    rotation {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "      w: 1.0\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 56.82294\n"
  "  data: 56.82294\n"
  "  data: 10.0\n"
  "}\n"
  "linear_damping: 0.0\n"
  "angular_damping: 0.0\n"
  "locked_rotation: false\n"
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
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/main/logo.atlas\"\n"
  "default_animation: \"logo\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
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

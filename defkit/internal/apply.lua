--- Handles targeting of which object the action should be applied to.
local M = {}

--- TARGETTING PROPERTIES
go.property("target_self", false)
go.property("target_other", false)
go.property("target_type", hash(""))

function M.apply(self, message_id, message)

end

return M

local M = {}

local dynamic_variables = {}

local empty_hash = hash("")

function M.add_object_by_type(object_type, game_object_id)

end

function M.remove_object(game_object_id)

end

function M.get_objects_by_type(object_type)

end

function M.set_variable(hashed, value)
	assert(hashed)
	assert(value)

	dynamic_variables[hashed] = value
end

function M.get_variable(hashed, default_value)
	assert(hashed)
	if hashed == empty_hash then
		return default_value
	end

	if dynamic_variables[hashed] then
		return dynamic_variables[hashed]
	elseif not default_value then
		error(string.format("Requested unknown variable %s", tostring(hashed)))
	end

	return default_value
end

return M

local M = {}

local dynamic_variables = {}

local objects_by_id = {}
local objects_by_type_id = {}

local EMPTY_HASH = hash("")

function M.add_object_by_type(object_type, game_object_id)
	assert(object_type)
	assert(game_object_id)

	local object = {
		id =		game_object_id,
		type =		object_type,
		url =	msg.url(),
	}

	objects_by_id[game_object_id] = object

	if not objects_by_type_id[object_type] then
		objects_by_type_id[object_type] = {}
	end
	objects_by_type_id[object_type][game_object_id] = object
end

function M.remove_object(game_object_id)
	assert(game_object_id)
	assert(objects_by_id[game_object_id], "Unknown game_object_id")

	local object = objects_by_id[game_object_id]

	objects_by_id[game_object_id] = nil
	objects_by_type_id[object.type][game_object_id] = nil
end

function M.get_objects_by_type(object_type)
	assert(object_type)

	return objects_by_type_id[object_type] or {}
end

function M.set_variable(hashed, value)
	assert(hashed)
	assert(value)

	dynamic_variables[hashed] = value
end

function M.get_variable(hashed, default_value)
	assert(hashed)

	if hashed == EMPTY_HASH then
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

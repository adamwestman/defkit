local M = {}

local dynamic_variables = {}

function M.add_object_by_type(object_type, game_object_id)
	
end

function M.remove_object(game_object_id)
	
end

function M.get_objects_by_type(object_type)
	
end

function M.set_variable(hash, value)
	assert(hash)
	assert(value)
	
	dynamic_variables[hash] = value
end

function M.get_variable(hash, default_value)
	assert(hash)
	
	if dynamic_variables[hash] then
		return dynamic_variables[hash]
	elseif not default_value then
		error(string.format("Requested unknown variable %s", tostring(hash)))
	end
	
	return default_value
end

return M
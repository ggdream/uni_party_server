if redis.call("LLEN", KEYS[1]) == ARGV[1] then
    return 0
else
    redis.call("LPUSH", KEYS[1], ARGV[2])
    return 1
end

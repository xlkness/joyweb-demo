import localCache from "@/stores/localCache";

class ExpireCache {
    // 添加
    setCache(key: string, value: any, expire: number) {
        let obj = {
            data: value,
            time: Date.now(),
            expire: expire * 1000,
        }
        const realKey = key  + ".expire"
        localCache.setCache(realKey, obj)
        // window.localStorage.setItem(key, JSON.stringify(obj))
    }
    // 查找
    getCache(key: string) {
        const realKey = key  + ".expire"
        let val = localCache.getCache(realKey)
        if (!val) {
            return val
        } else {
            if (Date.now() - val.time > val.expire) {
                localCache.deleteCache(realKey)
                return null
            }
            return val.data
        }
    }
    // 删除
    deleteCache(key: string) {
        const realKey = key  + ".expire"
        localCache.deleteCache(realKey)
    }
    // 清理
    clearCache() {
        // localCache.clearCache()
    }
}

export default new ExpireCache()
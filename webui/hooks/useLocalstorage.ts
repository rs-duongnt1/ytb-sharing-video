import { useState } from "react";

export const useLocalStorage = (key: string) => {
    const val = useState<string | null>(localStorage.getItem(key));
    const set = <T>(val: T) => {
        localStorage.setItem(key, val as string)
    }
    const remove = (key: string) => {
        localStorage.removeItem(key);
    }
    return {
        val,
        set,
        remove,
    }
}
import { DefaultValue } from "recoil";
// 定义一个通用的本地存储效果函数
export const localStorageEffect = <T>(key: string) => ({ setSelf, onSet }: { setSelf: (value: T | DefaultValue) => void, onSet: (handler: (value: T | DefaultValue) => void) => void }) => {
    // 在初始化时，从LocalStorage加载并设置Recoil状态
    const savedValue = localStorage.getItem(key);
    if (savedValue != null) {
        setSelf(JSON.parse(savedValue));
    }

    // 监听Recoil状态的变化，将新状态保存到LocalStorage
    onSet(newValue => {
        if (newValue instanceof DefaultValue) {
            localStorage.removeItem(key); // 如果值是 DefaultValue，则从本地存储中删除
        } else {
            localStorage.setItem(key, JSON.stringify(newValue)); // 否则保存新值到本地存储
        }
    });
};
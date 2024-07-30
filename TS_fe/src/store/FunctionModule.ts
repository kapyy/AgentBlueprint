import { atom, RecoilState } from "recoil";
import { localStorageEffect } from "../hooks/localStorage";



export type DataType = {
    id: number
    name:string
    data:any
};


// 定义一个通用的本地存储效果函数

  
  // 创建带有LocalStorage effect的Recoil atom
// 定义一个带有本地存储效果的Recoil状态(atom)
export const funListState: RecoilState<DataType[]> = atom({
    key: 'funListState',
    default: [], // 默认状态为空数组
    effects_UNSTABLE: [
        localStorageEffect<DataType[]>('funListState'), // 使用 DataType[] 类型的例子
    ],
});
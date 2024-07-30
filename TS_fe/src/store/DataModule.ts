import { atom, RecoilState } from "recoil";
import { localStorageEffect } from "../hooks/localStorage";



export type DataType = {
    id: number
    name:string
    data:any
};


export const dataListState: RecoilState<DataType[]> = atom({
    key: 'dataListState',
    default: [], // 默认状态为空数组
    effects_UNSTABLE: [
        localStorageEffect<DataType[]>('dataListState'), // 使用 DataType[] 类型的例子
    ],
});
import { atom, RecoilState } from "recoil";
import { localStorageEffect } from "../hooks/localStorage";



export type Node = {
    id: number
    name:string
    data:number,//这里其实应该是存id 其实是一个函数节点或者一个数据节点
    childNode: Node
};


// 这里其实的结构其实一个森林结构 ，需要对任意树的任意节点修改
export const treeListState: RecoilState<Node[]> = atom({
    key: 'treeListState',
    default: [], // 
    effects_UNSTABLE: [
        localStorageEffect<Node[]>('treeListState'), // 使用 DataType[] 类型的例子
    ],
});



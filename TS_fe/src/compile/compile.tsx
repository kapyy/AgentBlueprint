import { useRecoilState } from "recoil";
import { dataListState } from "../store/DataModule";
import { funListState } from "../store/FunctionModule";
import { treeListState } from "../store/TreeModule";

// store


// 
export function uesCompileNode(index:number){
    const [treeList,] = useRecoilState(treeListState)
    const [funList,] = useRecoilState(funListState)
    const [dataList,] = useRecoilState(dataListState)
    treeList[index]//当前需要编译的tree数据
    let nodes: never[]=[] 
    let edges: never[]=[]
    // 编译过程 
    return [nodes, edges]

}
//
function compileFile(){
    const [treeList,] = useRecoilState(treeListState)


}
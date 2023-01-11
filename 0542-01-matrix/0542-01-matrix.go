func updateMatrix(mat [][]int) [][]int {
    // BFS to update level 1: 10, level 2: 20; leve 3: 30    
    var cur []pair // get all [0]        
    for i:=0;i<len(mat);i++{
        for j:=0;j<len(mat[0]);j++{
            if mat[i][j]==0{cur=append(cur, pair{row: i,col: j,}) }            
        }
    }
    
    curlevel := 10
    rowD, colD:=[]int{1,-1, 0 , 0}, []int{0,0, 1 , -1}
    
    for len(cur)>0{
        var next []pair        
        for _, p := range cur{            
            for i:=0;i<4;i++{
                r, c := p.row+rowD[i], p.col+colD[i]                
                if r<0 ||c<0 || r==len(mat)|| c==len(mat[0]) ||mat[r][c]!=1{
                    continue
                }
                mat[r][c]=curlevel
                next=append(next, pair{row:r,col:c,})
            }
        }        
        cur=next
        curlevel+=10        
    }    
    // change the distance 
    for i:=0;i<len(mat);i++{
        for j:=0;j<len(mat[0]);j++{ mat[i][j]=mat[i][j]/10}
    } 
    return mat
}

type pair struct{
    row int
    col int
}
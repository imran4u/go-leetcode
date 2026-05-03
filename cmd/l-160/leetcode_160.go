package main

func main() {

}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
 // O(N*M) ~ O(N^2)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    // dummy := ListNode{}
    // temp := dummy
    pA := headA
    pB := headB
    for pA != nil {
        pB = headB
        for pB != nil{
            if pA == pB{
                return pA
            }
            pB = pB.Next
        }
        pA = pA.Next
    }
    return nil
}
*/

/*
//O(N) & O(N) memory
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    
    if headA == nil || headB == nil {
        return nil
    }
    pA := headA
    pB := headB
    addressMap := make(map[*ListNode]struct{})
    for pA != nil {
       addressMap[pA] = struct{}{}
       pA = pA.Next
    }
    for pB != nil {
        _, ok := addressMap[pB]
        if ok {
            return pB
        }
      pB = pB.Next
    }

    return nil
}
*/

/*
// conept: last element count after intersection is same, so their could be difference in intial elment only.
// remove the extra intial (after removal both have same no. of element before intersection) and then compare till the intersection.
// O(m+n) ~ O(N)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    
    if headA == nil || headB == nil {
        return nil
    }
    pA := headA
    pB := headB
    lengthA :=0
    lengthB := 0
    //calculate lenght A 
    for pA != nil {
        pA = pA.Next
        lengthA++
    }
    
    //calculate lenght B
    for pB != nil {
        pB = pB.Next
        lengthB++
    }
    
    // reset the pointer
    pA = headA
    pB = headB

    for lengthA >lengthB {
        lengthA--
        pA = pA.Next
    }

    for lengthB >lengthA {
        lengthB--
        pB = pB.Next
    }

    for pA != pB {
        pA= pA.Next
        pB = pB.Next
    }


    return pA
}
*/


// O(m+n) ~ O(N)
//Note : if you will run both the pointer in both list, you will must have intersection value or nil at last
// A1,A2,C1->nil, 
// B1,C1->nil
// so first loop pointer := A1, A2, C1, nil, B1, C1, nil
// so end loop pointer   := B1, C1, nil, A1, A2, C1, nil 
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    
    if headA == nil || headB == nil {
        return nil
    }
    pA := headA
    pB := headB
   
   // this condition will fail

    for pA != pB {
       if pA == nil {
        pA = headB
       }else {
        pA = pA.Next
       }
        if pB == nil {
            pB = headA
        }else {
            pB = pB.Next
        }

    }
    

    return pA
}
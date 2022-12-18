/**
 * See go 'Flower' struct
 * - colorPetal: Hex color of petals (without #)
 * - id: Unique unsigned ID
 * - parentID: Array of parents' ids, if they exist
 * - childID: Array of childrens' ids, if they exist
 */
export interface FlowerProps {
    colorPetal: string,
    id: number,
    parentID: number[] | null
    childID : number[] | null
}
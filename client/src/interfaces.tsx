/**
 * See Go 'Flower' struct
 * - id: Unique unsigned ID
 * - colorPetal: Hex color of petals (without #)
*/
export interface FlowerProps {
    id: number;
    colorPetal: string;
}

/**
 * See Go 'Box' struct
 * - id: Unique unsigned ID
 * - flower: This box's flower, if it exists
 */
export interface BoxProps {
    id: number;
    flower: FlowerProps | null;
}
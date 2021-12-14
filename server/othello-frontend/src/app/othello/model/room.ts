export interface Player {
  id: string,
  name: string,
  isApprove: boolean,
}

export interface Room {
  id: string,
  owner: Player,
  players: Player[],
  created: Date,
}

export interface Turn {
  blackId: string,
  whiteId: string,
}

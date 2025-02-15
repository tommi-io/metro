package structs


type Metro struct{
  fermate  *Fermata
  archi    *Arco
}

type Fermata struct{
  nome      string
  
  archi    *ArcPointer
  next     *Fermata
}

type ArcPointer struct{
  arco     *Arco
  next     *ArcPointer
}

type Arco struct{
  partenza *Fermata
  arrivo   *Fermata
  costo     int //tempo di percorrenza
  next     *Arco //è comunque una lista di archi
}


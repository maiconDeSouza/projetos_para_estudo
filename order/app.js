class Order {
    constructor(){
        this.id = crypto.randomUUID()
        
    }

    orderID(){
        return this.id
    }
}

class Item{
    constructor(name, price, quantity){
        this.name = name
        this.price = price
        this.quantity = quantity
    }

    total(){
        return this.price * this.quantity
    }
}
import { Sequelize } from "sequelize-typescript";
import Order from "../../../../domain/checkout/entity/order";
import OrderItem from "../../../../domain/checkout/entity/order_item";
import Customer from "../../../../domain/customer/entity/customer";
import Address from "../../../../domain/customer/value-object/address";
import Product from "../../../../domain/product/entity/product";
import CustomerModel from "../../../customer/repository/sequelize/customer.model";
import CustomerRepository from "../../../customer/repository/sequelize/customer.repository";
import ProductModel from "../../../product/repository/sequelize/product.model";
import ProductRepository from "../../../product/repository/sequelize/product.repository";
import OrderItemModel from "./orm/order-item.model";
import OrderModel from "./orm/order.model";
import OrderRepository from "./order.repository";

describe("Order repository test", () => {
  let sequelize: Sequelize;

  beforeEach(async () => {
    sequelize = new Sequelize({
      dialect: "sqlite",
      storage: ":memory:",
      logging: false,
      sync: { force: true },
    });

    await sequelize.addModels([
      CustomerModel,
      OrderModel,
      OrderItemModel,
      ProductModel,
    ]);
    await sequelize.sync();
  });

  afterEach(async () => {
    await sequelize.close();
  });

  it("should create a new order", async () => {
    const customerRepository = new CustomerRepository();
    const customer = new Customer("123", "Customer 1");
    const address = new Address("Street 1", 1, "Zipcode 1", "City 1");
    customer.changeAddress(address);
    await customerRepository.create(customer);

    const productRepository = new ProductRepository();
    const product = new Product("123", "Product 1", 10);
    await productRepository.create(product);

    const orderItem = new OrderItem(
      "1",
      product.name,
      product.price,
      product.id,
      2
    );

    const order = new Order("123", "123", [orderItem]);

    const orderRepository = new OrderRepository();
    await orderRepository.create(order);

    const orderModel = await OrderModel.findOne({
      where: { id: order.id },
      include: ["items"],
    });

    expect(orderModel.toJSON()).toStrictEqual({
      id: "123",
      customer_id: "123",
      total: order.total(),
      items: [
        {
          id: orderItem.id,
          name: orderItem.name,
          price: orderItem.price,
          quantity: orderItem.quantity,
          order_id: "123",
          product_id: "123",
        },
      ],
    });
  });


  it("should update an order", async () => {

    //First Product
    const customerRepository = new CustomerRepository();
    const customer = new Customer("123", "Customer 1");
    const address = new Address("Street 1", 1, "Zipcode 1", "City 1");
    customer.changeAddress(address);
    await customerRepository.create(customer);

    const productRepository = new ProductRepository();
    const product = new Product("123", "Product 1", 10);
    await productRepository.create(product);

    const orderItem = new OrderItem(
      "1",
      product.name,
      product.price,
      product.id,
      2
    );

    const order = new Order("123", "123", [orderItem]);

    const orderRepository = new OrderRepository();
    await orderRepository.create(order);


    //Second Product
    const customerRepository2 = new CustomerRepository();
    const customer2 = new Customer("456", "Customer 2");
    const address2 = new Address("Street 2", 2, "Zipcode 2", "City 2");
    customer2.changeAddress(address2);
    await customerRepository2.create(customer2);

    const productRepository2 = new ProductRepository();
    const product2 = new Product("456", "Product 2", 20);
    await productRepository2.create(product2);

    const orderItem2 = new OrderItem(
      "2",
      product2.name,
      product2.price,
      product2.id,
      3
    );

    const order2 = new Order("123", "456", [orderItem2]);

    await orderRepository.update(order2);


    const orderModel = await OrderModel.findOne({
      where: { id: order2.id },
      include: ["items"],
    });

    expect(orderModel.toJSON()).toStrictEqual({
      id: "123",
      customer_id: "456",
      total: order2.total(),
      items: [
        {
          id: orderItem2.id,
          name: orderItem2.name,
          price: orderItem2.price,
          quantity: orderItem2.quantity,
          order_id: "123",
          product_id: "456",
        },
      ],
    });
  })

  it('should find an order', async () => {
    const customerRepository = new CustomerRepository();
    const customer = new Customer('2', 'Pierry');
    const address = new Address('nomeRua', 2, '212121221', 'Rio de Janeiro');
    customer.changeAddress(address);
    await customerRepository.create(customer);

    const productRepository = new ProductRepository();
    const product = new Product('1', 'Product 1', 100);
    await productRepository.create(product);

    const orderItem = new OrderItem(
      '1',
      product.name,
      product.price,
      product.id,
      2
    );

    const order = new Order('1', customer.id, [orderItem]);
    const orderRepository = new OrderRepository();
    await orderRepository.create(order);

    const foundOrder = await orderRepository.find(order.id);
    expect(foundOrder).toStrictEqual(order);
  });

  it('should find all orders', async () => {
    const customerRepository = new CustomerRepository();
    const customer = new Customer('1', 'Pierry');
    const address = new Address('Rua 1', 1, '12345678', 'Rio de Janeiro');
    customer.changeAddress(address);
    const productRepository = new ProductRepository();
    const product = new Product('1', 'Product 1', 100);
    const orderItem = new OrderItem(
      '1',
      product.name,
      product.price,
      product.id,
      2
    );
    const order = new Order('1', customer.id, [orderItem]);

    const orderRepository = new OrderRepository();
    await customerRepository.create(customer);
    await productRepository.create(product);
    await orderRepository.create(order);

    let foundOrders = await orderRepository.findAll();
    expect(foundOrders).toStrictEqual([order]);

    const customer2 = new Customer('2', 'José');
    const address2 = new Address('Rua 2', 4, '12305976', 'Rio de Janeiro');
    customer2.changeAddress(address2);
    const product2 = new Product('2', 'Product 2', 200);
    const orderItem2 = new OrderItem(
      '2',
      product2.name,
      product2.price,
      product2.id,
      2
    );
    const order2 = new Order('2', customer2.id, [orderItem2]);
    await customerRepository.create(customer2);
    await productRepository.create(product2);
    await orderRepository.create(order2);

    foundOrders = await orderRepository.findAll();
    expect(foundOrders).toStrictEqual([order, order2]);
  });
})
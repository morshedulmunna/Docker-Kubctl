const { PrismaClient } = require("@prisma/client");

let prisma = null;

// Initialize and return Prisma client
const getPrismaClient = () => {
   if (!prisma) {
      prisma = new PrismaClient();
      console.log("Prisma client connected");
   }
   return prisma;
};

// Disconnect Prisma client gracefully
const disconnectPrismaClient = async () => {
   if (prisma) {
      try {
         await prisma.$disconnect();
         prisma = null;
         console.log("Prisma client disconnected");
      } catch (err) {
         console.log("Prisma client disconnection failed");
      }
   }
};
const prismaInstance = getPrismaClient();

module.exports = {
   prisma: prismaInstance,
   disconnectPrismaClient,
};

using CouInjector;


if (args.Length > 0)
{
    
   if (Injection.Run(args[0]))
    {
        Console.WriteLine("0");
    }
    else
    {
        Console.WriteLine("1");
    }
}
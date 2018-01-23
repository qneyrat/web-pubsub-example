<?php
declare(strict_types=1);

namespace App\Command;

use App\Entity\Conversation;
use App\Entity\Message;
use App\Entity\User;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;
use Symfony\Component\Security\Core\Encoder\UserPasswordEncoderInterface;

class SetupCommand  extends Command
{
    /**
     * @var EntityManagerInterface
     */
    private $entityManager;

    /**
     * @var UserPasswordEncoderInterface
     */
    private $encoder;

    /**
     * CreateUserCommand constructor.
     * @param EntityManagerInterface $entityManager
     * @param UserPasswordEncoderInterface $encoder
     */
    public function __construct(EntityManagerInterface $entityManager, UserPasswordEncoderInterface $encoder)
    {
        parent::__construct();
        $this->entityManager = $entityManager;
        $this->encoder = $encoder;
    }

    protected function configure()
    {
        $this->setName('app:setup');
    }

    protected function execute(InputInterface $input, OutputInterface $output)
    {
        $user1 = new User();
        $user1->setUsername('test1');
        $user1->setPassword($this->encoder->encodePassword($user1, 'test1'));

        $this->entityManager->persist($user1);

        $user2 = new User();
        $user2->setUsername('test2');
        $user2->setPassword($this->encoder->encodePassword($user2, 'test2'));

        $this->entityManager->persist($user2);

        $conversation = new Conversation();
        $conversation->addUser($user1);
        $conversation->addUser($user2);

        $message = new Message();
        $message->setFrom($user1);
        $message->setBody('hello world!');
        $message->setConversation($conversation);

        $conversation->addMessage($message);

        $this->entityManager->persist($conversation);

        $this->entityManager->flush();
    }
}

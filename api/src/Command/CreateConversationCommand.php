<?php
declare(strict_types=1);

namespace App\Command;

use App\Entity\Conversation;
use App\Entity\Message;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;

class CreateConversationCommand  extends Command
{
    /**
     * @var EntityManagerInterface
     */
    private $entityManager;

    /**
     * CreateConversationCommand constructor.
     * @param EntityManagerInterface $entityManager
     */
    public function __construct(EntityManagerInterface $entityManager)
    {
        parent::__construct();
        $this->entityManager = $entityManager;
    }

    protected function configure()
    {
        $this->setName('app:create-conversation');
    }

    protected function execute(InputInterface $input, OutputInterface $output)
    {
        $userId = '5a5a13ea9e3773014d6c9f21';
        $conversation = new Conversation();
        $conversation->setUsers([$userId]);

        $message = new Message();
        $message->setFrom($userId);
        $message->setBody('hello world!');

        $conversation->setMessages([$message]);

        $this->entityManager->persist($conversation);
        $this->entityManager->flush();
    }
}
